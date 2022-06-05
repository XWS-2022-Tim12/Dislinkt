import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Job } from 'src/app/model/job';
import { JobService } from 'src/app/service/job.service';

@Component({
  selector: 'app-add-salary-to-job',
  templateUrl: './add-salary-to-job.component.html',
  styleUrls: ['./add-salary-to-job.component.css']
})
export class AddSalaryToJobComponent implements OnInit {
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

  addSalaryToJob() {
    let salary = document.getElementById('salary') as HTMLTextAreaElement;
    let type = document.getElementById('type') as HTMLSelectElement;
    
    if(salary.value === ''){
      this.showError = true;
      return;
    } else
      this.showError = false;

    if(type.value === 'Junior') {    
      this.job.juniorSalary.push(Number(salary.value))   
    } else {
      this.job.mediorSalary.push(Number(salary.value))
    }
    

    this.jobService.editJob(this.job).subscribe(ret => {
      if(ret)
        this.router.navigate(['/profile']);
    });
  }
}
