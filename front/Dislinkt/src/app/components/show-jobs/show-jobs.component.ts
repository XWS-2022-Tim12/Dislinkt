import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { JobService } from 'src/app/services/job.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-show-jobs',
  templateUrl: './show-jobs.component.html',
  styleUrls: ['./show-jobs.component.css']
})
export class ShowJobsComponent implements OnInit {

  allJobs: any;
  showError: Boolean = false;

  constructor(public router: Router, private jobService: JobService) { }

  ngOnInit(): void {
      this.jobService.getRecomendedJobs().subscribe(ret => {
        this.allJobs = ret;
      })
  }

}
