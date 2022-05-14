import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  user: User;
  emailExists: boolean;
  usernameExists: boolean;
  registerForm:any;
  constructor(private userService: UserService, public router: Router) {
    this.user = new User();
   }

  ngOnInit(): void {
    this.emailExists = false;
    this.usernameExists = false;
    this.registerForm = new FormGroup({
      "username": new FormControl(null, [Validators.required,Validators.pattern('[a-zA-Z0-9]{4,20}')]),    
      "password": new FormControl(null, [Validators.required,Validators.pattern('[a-zA-Z0-9]{8,20}')]),
      "firstName": new FormControl(null, [Validators.required,Validators.pattern('[A-ZŠĐČĆŽ]{1}[a-zšđčćž]+')]),
      "email": new FormControl(null, [Validators.required,Validators.email]),
      "mobileNumber": new FormControl(null, [Validators.required,Validators.pattern('[0-9]{6,14}')]),
      "gender": new FormControl(null, [Validators.required,Validators.pattern('[A-ZŠĐČĆŽ]{1}[a-zšđčćž]+')]),
      "birthDay": new FormControl(null, [Validators.required])
    });
  }

  submitData()
  {
    this.user.mobileNumber = this.user.mobileNumber.toString();

    let dateTimeUTC = new Date(this.user.birthDay+"T12:00:00Z");
    this.user.birthDay = dateTimeUTC;
      this.userService.register(this.user).subscribe(ret => {   
          this.router.navigate(['/']);
      },
      (error: HttpErrorResponse) => {
        this.emailExists = true;
        this.usernameExists = true;
      });
  }

  login(){
    this.router.navigate(['/']);
  }

  get username() {
    return this.registerForm.get('username');
  }
  get password() {
    return this.registerForm.get('password');
  }
  get firstName() {
    return this.registerForm.get('firstName');
  }
  get email() {
    return this.registerForm.get('email');
  }
  get mobileNumber() {
    return this.registerForm.get('mobileNumber');
  }
  get gender() {
    return this.registerForm.get('gender');
  }
  get birthDay() {
    return this.registerForm.get('birthDay');
  }

}
