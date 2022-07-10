import { UserService } from 'src/app/services/user.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { User } from 'src/app/model/user';
import { DomSanitizer } from '@angular/platform-browser';
import { Post } from 'src/app/model/post';
import { PostService } from 'src/app/services/post.service';
import { MessageService } from 'src/app/services/message.service';
import { NotificationService } from 'src/app/services/notification.service';
import { Notification } from 'src/app/model/notification';

@Component({
  selector: 'app-other-user-profile',
  templateUrl: './other-user-profile.component.html',
  styleUrls: ['./other-user-profile.component.css']
})
export class OtherUserProfileComponent implements OnInit {
  user: User;
  username: string;
  loggedUser: User;
  usersInInbox: any[] = [];
  userPosts: Array<Post>;
  comments: string[] = [];

  canFollow: boolean;
  following: boolean;
  followingRequest: boolean;
  inboxOpen: boolean = false;

  blockedNotifications: boolean;

  constructor(private route: ActivatedRoute, private postService: PostService, private userService: UserService, private notificationService: NotificationService, public sanitizer: DomSanitizer, private router: Router) { }

  ngOnInit(): void {
    this.userService.getUserByUsername(sessionStorage.getItem("username")).subscribe(user => {
      this.loggedUser = Object.values(user)[0];

      this.userService.getUsernamesInInbox(this.loggedUser.username).subscribe(usersInInbox => {
        this.usersInInbox = usersInInbox;
      });

      this.route.params.subscribe(params => {
        this.username = params['username'];

        this.postService.getUserPosts(this.username).subscribe(posts => {
          this.userPosts = posts;
          for (let p of this.userPosts){
            if (!p.likes)
              p.likes = 0;
            if (!p.dislikes) 
              p.dislikes = 0;
            this.comments.push("")
          }
        })

        this.userService.getUserByUsername(params['username']).subscribe(user => {
          this.user = Object.values(user)[0];
          if (this.loggedUser.username == this.username) {
            this.canFollow = false;
            this.following = false;
            this.followingRequest = false;
            this.blockedNotifications = false;
            return
          }
          if (sessionStorage.getItem("username") != null) {
            this.canFollow = true;
            this.following = false;
            this.followingRequest = false;
          }
          for (let u of this.user.followedByUsers) {
            if (u == sessionStorage.getItem("username")) {
              this.canFollow = false;
              this.following = true;
              this.followingRequest = false;
              break;
            }
          }
          for (let u of this.user.followingRequests) {
            if (u == sessionStorage.getItem("username")) {
              this.canFollow = false;
              this.following = false;
              this.followingRequest = true;
              break;
            }
          }
          for (let u of this.loggedUser.notificationOffUsers) {
            if (u == this.username) {
              this.blockedNotifications = true;
              break;
            }
          }
        });
      });
    });
  }

  followUser(): void {
    let userToSend = new User();
    userToSend.id = this.loggedUser.id;
    userToSend.username = this.user.username;
    this.userService.follow(userToSend).subscribe(ret => {
      let notification = new Notification();
      notification.sender = this.loggedUser.username;
      notification.receiver = this.user.username;
      notification.creationDate = new Date();
      notification.notificationType = "follow";
      notification.description = "User " + this.loggedUser.username + " wants to follow " + this.user.username + ".";
      notification.isRead = false;
      this.notificationService.addNewNotification(notification).subscribe(ret => {

      });
      window.location.reload();
    })
  }

  likePost(post: Post): void {
    this.postService.likePost(post).subscribe(
      ret => {
      let notification = new Notification();
      notification.sender = this.loggedUser.username;
      notification.receiver = post.username;
      notification.creationDate = new Date();
      notification.notificationType = "like";
      notification.description = "User " + this.loggedUser.username + " liked post from " + post.username + ".";
      notification.isRead = false;
      this.notificationService.addNewNotification(notification).subscribe(ret => {

      });
        window.location.reload();
      })
  }

  dislikePost(post: Post): void {
    this.postService.dislikePost(post).subscribe(ret => {
      let notification = new Notification();
      notification.sender = this.loggedUser.username;
      notification.receiver = post.username;
      notification.creationDate = new Date();
      notification.notificationType = "dislike";
      notification.description = "User " + this.loggedUser.username + " disliked post from " + post.username + ".";
      notification.isRead = false;
      this.notificationService.addNewNotification(notification).subscribe(ret => {

      });
      window.location.reload();
    })
  }

  commentFunc(post: Post, i: number): void {
    if (this.comments[i]) {
      if (post.comments == null)
        post.comments = []

      post.comments.push(this.comments[i])
      this.postService.commentPost(post).subscribe(ret => {
        let notification = new Notification();
        notification.sender = this.loggedUser.username;
        notification.receiver = post.username;
        notification.creationDate = new Date();
        notification.notificationType = "comment";
        notification.description = "User " + this.loggedUser.username + " commented post from " + post.username + ".";
        notification.isRead = false;
        this.notificationService.addNewNotification(notification).subscribe(ret => {

        });
      })
    }
  }

  message(username: string): void {
    this.router.navigate(['/message', username]);
  }

  inbox(): void {
    if (this.inboxOpen) {
      this.inboxOpen = false;
    } else {
      this.inboxOpen = true;
    }
  }
  
  blockNotifications(username: string) {
    let user = new User();
    user.id = this.loggedUser.id;
    user.username = username;
    this.userService.changeNotificationsUsers(user).subscribe(ret => {
      window.location.reload();
    })
  }
}
