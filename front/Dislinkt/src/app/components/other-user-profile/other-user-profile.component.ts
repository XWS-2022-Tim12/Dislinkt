import { UserService } from 'src/app/services/user.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { User } from 'src/app/model/user';

@Component({
  selector: 'app-other-user-profile',
  templateUrl: './other-user-profile.component.html',
  styleUrls: ['./other-user-profile.component.css']
})
export class OtherUserProfileComponent implements OnInit {
  user: User;
  username: string;
  loggedUser: User;

  canFollow: boolean = false;
  following: boolean = false;
  followingRequest: boolean = false;

  constructor(private route: ActivatedRoute, private userService: UserService) { }

  ngOnInit(): void {
    this.userService.getUserByUsername(sessionStorage.getItem("username")).subscribe(user => {
      this.loggedUser = Object.values(user)[0];
    });
    this.route.params.subscribe(params => {
      this.username = params['username'];
      this.userService.getUserByUsername(params['username']).subscribe(user => {
        this.user = Object.values(user)[0];
        if (sessionStorage.getItem("username") != null) {
          this.canFollow = true;
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
  }

  followUser(): void {
    let userToSend = new User();
    userToSend.id = this.loggedUser.id;
    userToSend.username = this.user.username;
    this.userService.follow(userToSend).subscribe(ret => {
      window.location.reload();
    })
  }

}
