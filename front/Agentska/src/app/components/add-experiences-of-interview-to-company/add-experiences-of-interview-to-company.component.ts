import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/app/model/company';
import { CompanyService } from 'src/app/service/company.service';

@Component({
  selector: 'app-add-experiences-of-interview-to-company',
  templateUrl: './add-experiences-of-interview-to-company.component.html',
  styleUrls: ['./add-experiences-of-interview-to-company.component.css']
})
export class AddExperiencesOfInterviewToCompanyComponent implements OnInit {
  company: Company;
  showError: Boolean = false;

  constructor(public router: Router, private companyService: CompanyService) { }

  ngOnInit(): void {
    this.companyService.getCompanyByName().subscribe(ret => {
      this.company = ret;
    })
  }

  backToUsersPage() {
    this.router.navigate(['/profile']);
  }

  addExperienceToCompany() {
    let experience = document.getElementById('experience') as HTMLTextAreaElement;
    let type = document.getElementById('type') as HTMLSelectElement;
    
    if(experience.value === ''){
      this.showError = true;
      return;
    } else
      this.showError = false;

    if(type.value === 'HR') {
      if(this.company.hrInterviews === null){
        this.company.hrInterviews = new Array<String>();
        this.company.hrInterviews.push(experience.value)
      } else {
        this.company.hrInterviews.push(experience.value)
      }
    } else {
      if(this.company.tehnicalInterviews === null){
        this.company.tehnicalInterviews = new Array<String>();
        this.company.tehnicalInterviews.push(experience.value)
      } else {
        this.company.tehnicalInterviews.push(experience.value)
      }
    }
    

    this.companyService.editCompany(this.company).subscribe(ret => {
      if(ret)
        this.router.navigate(['/profile']);
    });
  }
}
