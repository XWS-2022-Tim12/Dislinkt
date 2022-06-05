import { UserService } from 'src/app/services/user.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { User } from 'src/app/model/user';
import { DomSanitizer } from '@angular/platform-browser';
import { Post } from 'src/app/model/post';
import { PostService } from 'src/app/services/post.service';

@Component({
  selector: 'app-other-user-profile',
  templateUrl: './other-user-profile.component.html',
  styleUrls: ['./other-user-profile.component.css']
})
export class OtherUserProfileComponent implements OnInit {
  user: User;
  username: string;
  loggedUser: User;
  userPosts: Array<Post>;
  comments: string[] = [];

  canFollow: boolean;
  following: boolean;
  followingRequest: boolean;

  constructor(private route: ActivatedRoute,  private postService: PostService, private userService: UserService, public sanitizer: DomSanitizer) { }

  ngOnInit(): void {
    this.userService.getUserByUsername(sessionStorage.getItem("username")).subscribe(user => {
      this.loggedUser = Object.values(user)[0];
    });
    this.route.params.subscribe(params => {
      this.username = params['username'];
      this.userService.getUserByUsername(params['username']).subscribe(user => {
        this.user = Object.values(user)[0];
        if (this.loggedUser.username == this.username) {
          this.canFollow = false;
          this.following = false;
          this.followingRequest = false;
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
      })
    })
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
  }

  followUser(): void {
    let userToSend = new User();
    userToSend.id = this.loggedUser.id;
    userToSend.username = this.user.username;
    this.userService.follow(userToSend).subscribe(ret => {
      window.location.reload();
    })
  }

  likePost(post: Post): void {
    this.postService.likePost(post).subscribe(
      ret => {
        window.location.reload();
      })
  }

  dislikePost(post: Post): void {
    this.postService.dislikePost(post).subscribe(ret => {
      window.location.reload();
    })
  }

  commentFunc(post: Post, i: number): void {
    if (this.comments[i]) {
      if (post.comments == null)
        post.comments = []

      post.comments.push(this.comments[i])
      this.postService.commentPost(post).subscribe(ret => {

      })
    }
  }
}
