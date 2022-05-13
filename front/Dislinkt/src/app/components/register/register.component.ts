import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user';
import { UserService } from 'src/app/services/user.service';

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
      "firstname": new FormControl(null, [Validators.required,Validators.pattern('[A-ZŠĐČĆŽ]{1}[a-zšđčćž]+')]),
      "email": new FormControl(null, [Validators.required,Validators.email]),
      "mobileNumber": new FormControl(null, [Validators.required,Validators.pattern('[0-9]{6,14}')]),
      "gender": new FormControl(null, [Validators.required,Validators.pattern('[A-ZŠĐČĆŽ]{1}[a-zšđčćž]+')]),
      "birthDay": new FormControl(null, [Validators.required]),
      "biography": new FormControl(null, [Validators.required,Validators.pattern('[a-zšđčćžA-ZŠĐČĆŽ ]*')]),
      "experience": new FormControl(null, [Validators.required,Validators.pattern('[a-zšđčćžA-ZŠĐČĆŽ ]*')]),
      "education": new FormControl(null, [Validators.required]),
      "skills": new FormControl(null, [Validators.required,Validators.pattern('[a-zšđčćžA-ZŠĐČĆŽ ]*')]),
      "interests": new FormControl(null, [Validators.required,Validators.pattern('[a-zšđčćžA-ZŠĐČĆŽ ]*')]),
      "public": new FormControl(null, [Validators.required])
    });
  }

  submitData()
  {
    if(this.user.public.toString() == "true"){
      this.user.public = true;
    }
    else{
      this.user.public = false;
    }
    this.user.mobileNumber = this.user.mobileNumber.toString();

    let dateTimeUTC = new Date(this.user.birthDay+"T12:00:00Z");
    this.user.birthDay = dateTimeUTC;
      this.userService.register(this.user).subscribe(ret => {   
        if (Object.values(ret)[0] == "email exists"){
          this.emailExists = true;
        }
        else if (Object.values(ret)[0] == "username exists"){
          this.usernameExists = true;
        }
        else {
          this.router.navigate(['/']);
        } 
      },
      (error: HttpErrorResponse) => {
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
  get firstname() {
    return this.registerForm.get('firstname');
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
  get biography() {
    return this.registerForm.get('biography');
  }
  get experience() {
    return this.registerForm.get('experience');
  }
  get education() {
    return this.registerForm.get('education');
  }
  get skills() {
    return this.registerForm.get('skills');
  }
  get interests() {
    return this.registerForm.get('interests');
  }
  get public() {
    return this.registerForm.get('public');
  }
}

