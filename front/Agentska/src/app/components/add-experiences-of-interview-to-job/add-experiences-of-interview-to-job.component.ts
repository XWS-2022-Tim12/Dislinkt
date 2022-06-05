import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Job } from 'src/app/model/job';
import { JobService } from 'src/app/service/job.service';

@Component({
  selector: 'app-add-experiences-of-interview-to-job',
  templateUrl: './add-experiences-of-interview-to-job.component.html',
  styleUrls: ['./add-experiences-of-interview-to-job.component.css']
})
export class AddExperiencesOfInterviewToJobComponent implements OnInit {
  job: Job;
  showError: Boolean = false;

  constructor(public router: Router, private jobService: JobService) { }

  ngOnInit(): void {
    this.jobService.getJobById(sessionStorage.getItem('id')).subscribe(ret => {
      this.job = ret;
    })
  }

  backToUsersPage() {
    this.router.navigate(['/profile']);
  }

  addExperienceToJob() {
    let experience = document.getElementById('experience') as HTMLTextAreaElement;
    let type = document.getElementById('type') as HTMLSelectElement;
    
    if(experience.value === ''){
      this.showError = true;
      return;
    } else
      this.showError = false;

    if(type.value === 'HR') {
      this.job.hrInterviews.push(experience.value)
    } else {
      this.job.tehnicalInterviews.push(experience.value) 
    }
    

    this.jobService.editJob(this.job).subscribe(ret => {
      if(ret)
        this.router.navigate(['/profile']);
    });
  }
}
