import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Job } from 'src/app/model/job';
import { Role, User } from 'src/app/model/user';
import { JobService } from 'src/app/service/job.service';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-add-new-job',
  templateUrl: './add-new-job.component.html',
  styleUrls: ['./add-new-job.component.css']
})
export class AddNewJobComponent implements OnInit {
  job: Job = new Job();

loggedUser: User = new User();

  constructor(private userService: UserService, private jobService: JobService, private router: Router) { }

  ngOnInit(): void {
    if (sessionStorage.getItem('username') == null) {
      this.router.navigate(['']);
      alert('You are not logged in!');
    } else {
      this.userService.getLoggedUser().subscribe(user => {
        if (user.role == Role.agent_owner || user.role.toString() == "agent_owner") {
          this.loggedUser = user;
        } else {
          this.router.navigate(['/profile']);
          alert('You are not logged in as company owner!');
        }
        this.loggedUser = user;
      });
	  }
  }

  newJob() {
    if (sessionStorage.getItem('username') != null) {
      this.job.userId = this.loggedUser.id;
      this.job.creationDay = new Date();
      this.jobService.addNewJob(this.job).subscribe(ret => {
        this.router.navigate(['/profile']);
        alert('Job successfully created!');
      }, (error: HttpErrorResponse) => {
        this.router.navigate(['/profile']);
        })
    } else {
      this.router.navigate(['']);
      alert('You are not logged in!');
    }
  }
  
}
