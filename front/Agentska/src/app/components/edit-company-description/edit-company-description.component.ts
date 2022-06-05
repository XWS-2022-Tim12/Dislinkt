import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Company } from 'src/app/model/company';
import { Role, User } from 'src/app/model/user';
import { CompanyService } from 'src/app/service/company.service';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-edit-company-description',
  templateUrl: './edit-company-description.component.html',
  styleUrls: ['./edit-company-description.component.css']
})
export class EditCompanyDescriptionComponent implements OnInit {
  company: Company = new Company();

  constructor(private companyService: CompanyService, private userService: UserService, private router: Router, private route: ActivatedRoute) { }

  ngOnInit(): void {
    if (sessionStorage.getItem('username') == null) {
      this.router.navigate(['']);
      alert('You are not logged in!');
    } else {
      this.userService.getLoggedUser().subscribe(user => {
        if (user.role == Role.agent_owner || user.role.toString() == "agent_owner") {
          this.route.params.subscribe(params => {
            this.companyService.getAllCompanies().subscribe(companies => {
              let found = false;
              for (let c of companies) {
                if (c.name == params['name']) {
                  if (c.owner.username == sessionStorage.getItem('username')) {
                    this.company = c;
                    found = true;
                    break;
                  } else {
                    this.router.navigate(['/profile']);
                    alert('You are not owner of this company!');
                  }
                }
              }
              if (!found) {
                this.router.navigate(['/profile']);
                alert('Company does not exist!');
              }
            });
          })
        } else {
          this.router.navigate(['/profile']);
          alert('You are not logged in as owner!');
        }
      })
    }
  }

  editCompany() {
    if (sessionStorage.getItem('username') != null) {
      this.company.owner = new User();
      this.company.owner.username = sessionStorage.getItem('username');
      this.userService.changeCompanyDescription(this.company).subscribe(ret => {
        if (ret) {
          this.router.navigate(['/profile']);
          alert('Company description changed!');
        } else {
          alert('Error with editing of company description. Please try again.');
        }
      })
    } else {
      this.router.navigate(['']);
      alert('You are not logged in!');
    }
  }

}
