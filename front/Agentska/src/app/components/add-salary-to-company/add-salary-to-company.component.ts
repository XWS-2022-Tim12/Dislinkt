import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/app/model/company';
import { CompanyService } from 'src/app/service/company.service';

@Component({
  selector: 'app-add-salary-to-company',
  templateUrl: './add-salary-to-company.component.html',
  styleUrls: ['./add-salary-to-company.component.css']
})
export class AddSalaryToCompanyComponent implements OnInit {
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

  addSalaryToCompany() {
    let salary = document.getElementById('salary') as HTMLTextAreaElement;
    let type = document.getElementById('type') as HTMLSelectElement;
    
    if(salary.value === ''){
      this.showError = true;
      return;
    } else
      this.showError = false;

    if(type.value === 'Junior') {
      if(this.company.juniorSalary === null){
        this.company.juniorSalary = new Array<Number>();
        this.company.juniorSalary.push(Number(salary.value))
      } else {
        this.company.juniorSalary.push(Number(salary.value))
      }
    } else {
      if(this.company.mediorSalary === null){
        this.company.mediorSalary = new Array<Number>();
        this.company.mediorSalary.push(Number(salary.value))
      } else {
        this.company.mediorSalary.push(Number(salary.value))
      }
    }
    

    this.companyService.editCompany(this.company).subscribe(ret => {
      if(ret)
        this.router.navigate(['/profile']);
    });
  }
}
