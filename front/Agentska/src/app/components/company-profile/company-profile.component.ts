import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/app/model/company';
import { CompanyService } from 'src/app/service/company.service';

@Component({
  selector: 'app-company-profile',
  templateUrl: './company-profile.component.html',
  styleUrls: ['./company-profile.component.css']
})
export class CompanyProfileComponent implements OnInit {
  company: Company;

  constructor(public router: Router, private companyService: CompanyService) { }

  ngOnInit(): void {
    this.companyService.getCompanyByName().subscribe(ret => {
      this.company = ret;
    })
  }

  backTousersPage() {
    this.router.navigate(['/profile']);
  }
}
