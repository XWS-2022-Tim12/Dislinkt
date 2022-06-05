import { Router } from '@angular/router';
import { Company } from './../../model/company';
import { Component, OnInit } from '@angular/core';
import { UserService } from 'src/app/service/user.service';
import { User } from 'src/app/model/user';

@Component({
  selector: 'app-register-company',
  templateUrl: './register-company.component.html',
  styleUrls: ['./register-company.component.css']
})
export class RegisterCompanyComponent implements OnInit {
  company: Company = new Company();

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    if (sessionStorage.getItem('username') == null) {
      this.router.navigate(['']);
      alert('You are not logged in!');
    }
  }

  registerCompany() {
    if (sessionStorage.getItem('username') != null) {
      this.company.owner = new User();
      this.company.owner.username = sessionStorage.getItem('username');
      this.userService.registerCompany(this.company).subscribe(ret => {
        if (ret) {
          alert('Company registrated!');
        } else {
          alert('Error with registration of company. Please try again.');
        }
      })
    } else {
      this.router.navigate(['']);
      alert('You are not logged in!');
    }
  }

}
