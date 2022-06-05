import { Router } from '@angular/router';
import { UserService } from './../../service/user.service';
import { CompanyService } from './../../service/company.service';
import { Component, OnInit } from '@angular/core';
import { Company } from 'src/app/model/company';
import { Role } from 'src/app/model/user';

@Component({
  selector: 'app-register-company-requests',
  templateUrl: './register-company-requests.component.html',
  styleUrls: ['./register-company-requests.component.css']
})
export class RegisterCompanyRequestsComponent implements OnInit {
  areRequestsEmpty = true;
  companyRequests: Array<Company> = new Array<Company>();

  constructor(private companyService: CompanyService, private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    if (sessionStorage.getItem('username') == null) {
      this.router.navigate(['']);
      alert('You are not logged in!');
    } else {
      this.userService.getLoggedUser().subscribe(user => {
        if (user.role == Role.admin || user.role.toString() == "admin") {
          this.companyRequests = new Array<Company>();
          this.companyService.getAllCompanies().subscribe(companies => {
            for (let c of companies) {
              if (!c.approved) {
                this.companyRequests.push(c);
                this.areRequestsEmpty = false;
              }
            }
          })
        } else {
          this.router.navigate(['/profile']);
          alert('You are not logged in as admin!');
        }
      })
      
    }
  }

  acceptRequest(company: Company) {
    this.userService.acceptCompanyRegistrationRequest(company).subscribe(ret => {
      if (ret) {
        window.location.reload();
        alert("Accepted!");
      } else {
        alert('Error with accepting the company. Please try again.')
      }
    })
  }

}
