import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user';
import { CompanyService } from 'src/app/service/company.service';
import { JobService } from 'src/app/service/job.service';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-users-page',
  templateUrl: './users-page.component.html',
  styleUrls: ['./users-page.component.css']
})
export class UsersPageComponent implements OnInit {
  user: User;
  allJobs: any;
  showError: Boolean = false;

  constructor(public router: Router, private userService: UserService, private jobService: JobService) { }

  ngOnInit(): void {
    this.userService.getLoggedUser().subscribe(ret => {
      this.user = ret;
      this.jobService.getAllJobs().subscribe(ret => {
        this.allJobs = ret;
      })
    });
    
  }

  addComment(index: number) {
    let job = this.allJobs[index]
    if(job.userId === this.user.id){
      this.showError = true;
      return;
    }

    sessionStorage.setItem('id', job.id);
    this.router.navigate(['/add-comment']);
  }

  showCompanyProfile(index: number) {
    let job = this.allJobs[index]
    if(job.userId === this.user.id){
      this.showError = true;
      return;
    }
    sessionStorage.setItem('id', job.id);
    this.router.navigate(['/job-profile']);
  }

  addSalary(index: number) {
    let job = this.allJobs[index]
    if(job.userId === this.user.id){
      this.showError = true;
      return;
    }
    sessionStorage.setItem('id', job.id);
    this.router.navigate(['/add-salary']);
  }

  addExperience(index: number) {
    let job = this.allJobs[index]
    if(job.userId === this.user.id){
      this.showError = true;
      return;
    }
    sessionStorage.setItem('id', job.id);
    this.router.navigate(['/add-experience']);
  }

  addJob() {
    this.router.navigate(['/new-job']);
  }
}
