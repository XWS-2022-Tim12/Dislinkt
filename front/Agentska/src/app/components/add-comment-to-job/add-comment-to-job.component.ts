import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Job } from 'src/app/model/job';
import { JobService } from 'src/app/service/job.service';

@Component({
  selector: 'app-add-comment-to-job',
  templateUrl: './add-comment-to-job.component.html',
  styleUrls: ['./add-comment-to-job.component.css']
})
export class AddCommentToJobComponent implements OnInit {
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

  addCommentToJob() {
    let comment = document.getElementById('comment') as HTMLTextAreaElement;

    if(comment.value === ''){
      this.showError = true;
      return;
    } else
      this.showError = false;

    this.job.comments.push(comment.value)
    
    this.jobService.editJob(this.job).subscribe(ret => {
      if(ret)
        this.router.navigate(['/profile']);
    });
  }
}
