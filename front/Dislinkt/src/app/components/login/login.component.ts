import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  user: User;
  userNotFound: boolean;
  loginForm:any;
  constructor(private userService: UserService, public router: Router) {
    this.user = new User();
   }

  ngOnInit(): void {
    this.userNotFound = false;
    this.loginForm = new FormGroup({
      "username": new FormControl(null, [Validators.required,Validators.pattern('[a-zA-Z0-9]{4,20}')]),    
      "password": new FormControl(null, [Validators.required,Validators.pattern('[a-zA-Z0-9]{8,20}')])
    });
  }

  submitData()
  {
      this.userService.login(this.user).subscribe(ret => { 
          sessionStorage.setItem("username",this.user.username);  
          this.router.navigate(['/profile']);
      },
      (error: HttpErrorResponse) => {
        this.userNotFound = true;
      });
  }

  register(){
    this.router.navigate(['/register']);
  }

  get username() {
    return this.loginForm.get('username');
  }
  get password() {
    return this.loginForm.get('password');
  }

}
