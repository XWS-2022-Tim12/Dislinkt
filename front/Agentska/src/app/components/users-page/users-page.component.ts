import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user';
import { CompanyService } from 'src/app/service/company.service';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-users-page',
  templateUrl: './users-page.component.html',
  styleUrls: ['./users-page.component.css']
})
export class UsersPageComponent implements OnInit {
  user: User;
  allCompanies: any;

  constructor(public router: Router, private userService: UserService, private companyService: CompanyService) { }

  ngOnInit(): void {
    this.userService.getLoggedUser().subscribe(ret => {
      this.user = ret;
    });
    this.companyService.getAllCompanies().subscribe(ret => {
      this.allCompanies = ret;
    })
  }

  addComment(index: number) {
    let company = this.allCompanies[index]
    sessionStorage.setItem('name', company.name);
    this.router.navigate(['/add-comment']);
  }

  showCompanyProfile(index: number) {
    let company = this.allCompanies[index]
    sessionStorage.setItem('name', company.name);
    this.router.navigate(['/company-profile']);
  }

  addSalary(index: number) {
    let company = this.allCompanies[index]
    sessionStorage.setItem('name', company.name);
    this.router.navigate(['/add-salary']);
  }

  addExperience(index: number) {
    let company = this.allCompanies[index]
    sessionStorage.setItem('name', company.name);
    this.router.navigate(['/add-experience']);
  }
}
