import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/model/user';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-following-requests',
  templateUrl: './following-requests.component.html',
  styleUrls: ['./following-requests.component.css']
})
export class FollowingRequestsComponent implements OnInit {
  isUserLoggedIn: boolean = false;
  loggedUser: User;

  areRequestsEmpty: boolean = true;
  users: Array<string>;

  constructor(private router: Router, private userService: UserService) { }

  ngOnInit(): void {
    this.users = new Array<string>();
    this.userService.getUserByUsername(sessionStorage.getItem("username")).subscribe(user => {
      this.loggedUser = Object.values(user)[0];
      this.isUserLoggedIn = true;
      for (let username of this.loggedUser.followingRequests) {
        this.areRequestsEmpty = false;
        this.users.push(username);
      }
    }, () => {
      this.router.navigate(['/login']);
    });
  }

  acceptRequest(username: string) {
    this.userService.getUserByUsername(username).subscribe(user => {
      let userToSend = new User();
      userToSend.id = Object.values(user)[0].id;
      userToSend.username = this.loggedUser.username;
      this.userService.acceptFollowingRequest(userToSend).subscribe(ret => {
        window.location.reload();
        alert('Accept status: ' + ret);
      });
    });
  }

  rejectRequest(username: string) {
    this.userService.getUserByUsername(username).subscribe(user => {
      let userToSend = new User();
      userToSend.id = Object.values(user)[0].id;
      userToSend.username = this.loggedUser.username;
      this.userService.rejectFollowingRequest(userToSend).subscribe(ret => {
        window.location.reload();
        alert('Reject status: ' + ret);
      });
    });
  }

}
