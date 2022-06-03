import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-home-page-guest',
  templateUrl: './home-page-guest.component.html',
  styleUrls: ['./home-page-guest.component.css']
})
export class HomePageGuestComponent implements OnInit {
  myControl = new FormControl();
  options: string[] = [];
  publicUsernames: string[] = [];
  searchText: string = "";
  userNotFound: boolean = false;

  constructor(private userService: UserService, public router: Router) { }

  ngOnInit(): void {
    this.userService.getAllPublicUsers().subscribe(ret => {
      let users = Object.values(ret)[0]
      for (let u of users){
        this.publicUsernames.push(u.username);
      }
    })
  }

  getUsersByUsername() {
    if (this.searchText == "")
      this.options = [];
    else
      this.options = this.publicUsernames.filter((username) => username.startsWith(this.searchText));
  }

  goToProfile() {
    this.userService.getPublicUserByUsername(this.searchText).subscribe(
      ret => {
        this.userNotFound = false;
        this.router.navigate(['/user/' + this.searchText]);
      },
      err => {this.userNotFound = true;})
  }

  register(){
    this.router.navigate(['/register']);
  }

  login(){
    this.router.navigate(['/login']);
  }
}
