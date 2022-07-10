import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/model/user';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-show-users',
  templateUrl: './show-users.component.html',
  styleUrls: ['./show-users.component.css']
})
export class ShowUsersComponent implements OnInit {
  allUsers: any;
  loggedUser: User;

  constructor(private userService: UserService) { }

  ngOnInit(): void {
    this.userService.getAll().subscribe(ret => {
      this.allUsers = Object.values(ret)[0]
      for(let u of this.allUsers){
        if(u.username === sessionStorage.getItem("username")){
          this.loggedUser = u;
        }
      }
      this.checkUsers()
    })
  }

  checkUsers() {
    for(let user of this.allUsers) {
      for(let blockedUser of user.blockedUsers) {
        if(blockedUser === this.loggedUser.username) {
          this.allUsers.splice(this.allUsers.indexOf(user), 1);
        }
      }
    }
  }

  isUserFollowing(username: string): Boolean {
    for(let followingUser of this.loggedUser.followingUsers) {
      if(followingUser === username) {
        return true
      } 
    }
    return false
  }

  isUserBlocked(username: string): Boolean {
    for(let blockedUser of this.loggedUser.blockedUsers) {
      if(blockedUser === username) {
        return true
      } 
    }
    return false
  }

  blockUser(username: string) {
    let following = false
    for(let followingUser of this.loggedUser.followingUsers) {
      if(followingUser === username) {
        this.loggedUser.followingUsers.splice(this.loggedUser.followingUsers.indexOf(followingUser), 1);
        following = true
      } 
    }
    this.loggedUser.blockedUsers.push(username)
    this.userService.blockUser(this.loggedUser).subscribe(ret => {
      if(following){
        let blockedUser = new User()
        for(let user of this.allUsers) {
          if(user.username === username)
            blockedUser = user
        }
        blockedUser.followedByUsers.splice(blockedUser.followedByUsers.indexOf(sessionStorage.getItem("username")), 1);
        this.userService.editAll(blockedUser).subscribe(ret => {

        })
      }
    })
  }

  unblockUser(username: string) {
    this.loggedUser.blockedUsers.splice(this.loggedUser.blockedUsers.indexOf(username), 1);
    this.userService.editAll(this.loggedUser).subscribe(ret => {
    })
  }
}
