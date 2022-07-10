import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user';
import { NotificationService } from 'src/app/services/notification.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  user: User;
  wrongPassword: boolean;
  usernameExists: boolean;
  profileForm:any;
  date: string;
  searchText: string;
  users: any;

  constructor(private userService: UserService, private notificationService: NotificationService, public router: Router) {
   }

  ngOnInit(): void {
    this.userService.getAll().subscribe(ret => {
      this.users = Object.values(ret)[0]
      for (let u of this.users){
        if(u.username == sessionStorage.getItem("username")){
          this.user = u;
          this.user.password = "";
          this.user.birthDay = new Date(this.user.birthDay);
          this.date =this.user.birthDay.toISOString().split("T")[0];

          this.notificationService.getNotificationsByReceiver(u.username).subscribe(nots => {
            let notifications = nots['notifications'];
            for (let not of notifications) {
              if (this.user.notificationOffUsers.includes(not.sender)) {
                if (not.notificationType != 'message') {
                  continue;
                }
              } else if (this.user.notificationOffMessages.includes(not.sender)) {
                if (not.notificationType == 'message') {
                  continue;
                }
              } 
              if (!not.isRead) {
                alert(not.creationDate + ': ' + not.description);
                not.isRead = true;
                this.notificationService.editNotification(not).subscribe(ret => {

                });
              }
            }
          })

        }
      }
    })
    this.wrongPassword = false;
    this.usernameExists = false;
    this.profileForm = new FormGroup({
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

  searchUser() {
    for(let user of this.users) {
      if(user.username === this.searchText){
        for(let block of user.blockedUsers) {
          if(block === this.user.username) {
            alert('Unknown user!')
            return;
          }
        }
      }
    }
    for(let user of this.user.blockedUsers) {
      if(user === this.searchText){
        alert('This user is blocked!')
        return;
      }
    }
    this.router.navigate(['/user/profile/' + this.searchText]);
  }

  editBasic()
  {
    if(this.user.public.toString() == "true"){
      this.user.public = true;
    }
    else{
      this.user.public = false;
    }
    this.user.mobileNumber = this.user.mobileNumber.toString();

    let dateTimeUTC = new Date(this.date+"T12:00:00Z");
    this.user.birthDay = dateTimeUTC;
      this.userService.editBasic(this.user).subscribe(ret => {   
        if (ret == "wrong password"){
          this.wrongPassword = true;
        }
        else if (ret == "username exists"){
          this.usernameExists = true;
        }
        else{
          sessionStorage.setItem("username", this.user.username);
          alert("successful");
        }
      },
      (error: HttpErrorResponse) => {
      });
  }

  editAdvanced()
  {
    if(this.user.public.toString() == "true"){
      this.user.public = true;
    }
    else{
      this.user.public = false;
    }
    let what = 0;
    if(this.user.education == "PrimaryEducation"){
      this.user.education = "Primary education"
      what = 1;
    }
    if(this.user.education == "LowerSecondaryEducation"){
      this.user.education = "Lower secondary education"
      what = 2;
    }
    if(this.user.education == "UpperSecondaryEducation"){
      this.user.education = "Upper secondary education"
      what = 3;
    }
    this.user.mobileNumber = this.user.mobileNumber.toString();

    let dateTimeUTC = new Date(this.date+"T12:00:00Z");
    this.user.birthDay = dateTimeUTC;
      this.userService.editAdvanced(this.user).subscribe(ret => {   
        if (ret == "wrong password"){
          this.wrongPassword = true;
        }
        else if (ret == "username exists"){
          this.usernameExists = true;
        }
        else{
          alert("successful");
          if(what == 1){
            this.user.education = "PrimaryEducation"
          }
          if(what == 2){
            this.user.education = "LowerSecondaryEducation"
          }
          if(what == 3){
            this.user.education = "UpperSecondaryEducation"
          }
        }
      },
      (error: HttpErrorResponse) => {
      });
  }

  editPersonal()
  {
    if(this.user.public.toString() == "true"){
      this.user.public = true;
    }
    else{
      this.user.public = false;
    }
    this.user.mobileNumber = this.user.mobileNumber.toString();

    let dateTimeUTC = new Date(this.date+"T12:00:00Z");
    this.user.birthDay = dateTimeUTC;
      this.userService.editPersonal(this.user).subscribe(ret => {   
        if (ret == "wrong password"){
          this.wrongPassword = true;
        }
        else if (ret == "username exists"){
          this.usernameExists = true;
        }
        else{
          alert("successful");
        }
      },
      (error: HttpErrorResponse) => {
      });
  }

  editAll()
  {
    if(this.user.public.toString() == "true"){
      this.user.public = true;
    }
    else{
      this.user.public = false;
    }
    let what = 0;
    this.user.mobileNumber = this.user.mobileNumber.toString();
    if(this.user.education == "PrimaryEducation"){
      this.user.education = "Primary education"
      what = 1;
    }
    if(this.user.education == "LowerSecondaryEducation"){
      this.user.education = "Lower secondary education"
      what = 2;
    }
    if(this.user.education == "UpperSecondaryEducation"){
      this.user.education = "Upper secondary education"
      what = 3;
    }

    let dateTimeUTC = new Date(this.date+"T12:00:00Z");
    this.user.birthDay = dateTimeUTC;
      this.userService.editAll(this.user).subscribe(ret => {   
        if (ret == "wrong password"){
          this.wrongPassword = true;
        }
        else if (ret == "username exists"){
          this.usernameExists = true;
        }
        else{
        sessionStorage.setItem("username", this.user.username);
        alert("successful");
        if(what == 1){
          this.user.education = "PrimaryEducation"
        }
        if(what == 2){
          this.user.education = "LowerSecondaryEducation"
        }
        if(what == 3){
          this.user.education = "UpperSecondaryEducation"
        }
        }
      },
      (error: HttpErrorResponse) => {
      });
  }

  showUsers() {
    this.router.navigate(['/showUsers']);
  }

  showSuggestions() {
    this.router.navigate(['/followingSuggestions']);
  }

  get username() {
    return this.profileForm.get('username');
  }
  get password() {
    return this.profileForm.get('password');
  }
  get firstname() {
    return this.profileForm.get('firstname');
  }
  get email() {
    return this.profileForm.get('email');
  }
  get mobileNumber() {
    return this.profileForm.get('mobileNumber');
  }
  get gender() {
    return this.profileForm.get('gender');
  }
  get birthDay() {
    return this.profileForm.get('birthDay');
  }
  get biography() {
    return this.profileForm.get('biography');
  }
  get experience() {
    return this.profileForm.get('experience');
  }
  get education() {
    return this.profileForm.get('education');
  }
  get skills() {
    return this.profileForm.get('skills');
  }
  get interests() {
    return this.profileForm.get('interests');
  }
  get public() {
    return this.profileForm.get('public');
  }

}
