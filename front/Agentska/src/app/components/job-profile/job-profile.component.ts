import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Job } from 'src/app/model/job';
import { JobService } from 'src/app/service/job.service';

@Component({
  selector: 'app-job-profile',
  templateUrl: './job-profile.component.html',
  styleUrls: ['./job-profile.component.css']
})
export class JobProfileComponent implements OnInit {
  job: Job;

  constructor(public router: Router, private jobService: JobService) { }

  ngOnInit(): void {
    this.jobService.getJobById(sessionStorage.getItem('id')).subscribe(ret => {
      this.job = ret;
    })
  }

  backTousersPage() {
    this.router.navigate(['/profile']);
  }
}
