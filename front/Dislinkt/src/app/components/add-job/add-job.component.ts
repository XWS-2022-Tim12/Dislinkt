import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Job } from 'src/app/model/job';
import { User } from 'src/app/model/user';
import { JobService } from 'src/app/services/job.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-add-job',
  templateUrl: './add-job.component.html',
  styleUrls: ['./add-job.component.css']
})
export class AddJobComponent implements OnInit {

  job: Job = new Job();
  constructor(private userService: UserService, private jobService: JobService, private router: Router) { }
  ngOnInit(): void {
  }

  newJob() {
    if (sessionStorage.getItem('username') != null) {
      this.userService.getUserByUsername(sessionStorage.getItem('username')).subscribe(user => {
        var id = Object.values(user)[0].id;
        this.job.userId = id;
        this.job.creationDay = new Date();
        this.jobService.addNewJob(this.job).subscribe(ret => {
        this.router.navigate(['/profile']);
        alert('Job successfully created!');
      }, (error: HttpErrorResponse) => {
        this.router.navigate(['/profile']);
        })
    })} else {
      this.router.navigate(['']);
      alert('You are not logged in!');
    }
      }

}
