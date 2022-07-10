import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/user.service';
import { User } from 'src/app/model/user';

@Component({
  selector: 'app-show-suggestions-for-following',
  templateUrl: './show-suggestions-for-following.component.html',
  styleUrls: ['./show-suggestions-for-following.component.css']
})
export class ShowSuggestionsForFollowingComponent implements OnInit {
  suggestedUsers: any;
  loggedUser: User;

  constructor(private userService: UserService, public router: Router) { }

  ngOnInit(): void {
    this.userService.getSuggestions().subscribe(ret => {
      this.suggestedUsers = ret;
    })
    this.userService.getUserByUsername(sessionStorage.getItem("username")).subscribe(user => {
      this.loggedUser = Object.values(user)[0];
    });
  }

  openProfile(index: number) {
      let user = this.suggestedUsers[index]
      this.userService.getPublicUserByUsername(user.username).subscribe(
        ret => {
          this.router.navigate(['/user/profile', user.username]);
        })
  }

  followUser(index: number) {
      let user = this.suggestedUsers[index]
      let userToSend = new User()
      userToSend.id = this.loggedUser.id;
      userToSend.username = user.username;
      this.userService.follow(userToSend).subscribe(ret => {
        window.location.reload();
      })
  }

  isUserFollowing(username: string): Boolean {
    for(let user of this.loggedUser.followingUsers) {
      if(user === username) {
        for(let suggestedUser of this.suggestedUsers) {
          if(suggestedUser.username === user) {
            this.suggestedUsers.splice(this.suggestedUsers.indexOf(suggestedUser), 1)
          }
        }

        return true
      }
    }

    return false
  }
}
