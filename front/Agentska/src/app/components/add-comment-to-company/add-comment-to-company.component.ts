import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/app/model/company';
import { CompanyService } from 'src/app/service/company.service';

@Component({
  selector: 'app-add-comment-to-company',
  templateUrl: './add-comment-to-company.component.html',
  styleUrls: ['./add-comment-to-company.component.css']
})
export class AddCommentToCompanyComponent implements OnInit {
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

  addCommentToCompany() {
    let comment = document.getElementById('comment') as HTMLTextAreaElement;

    if(comment.value === ''){
      this.showError = true;
      return;
    } else
      this.showError = false;

    if(this.company.comments === null){
      this.company.comments = new Array<String>();
      this.company.comments.push(comment.value)
    } else {
      this.company.comments.push(comment.value)
    }

    this.companyService.addCommentToCompany(this.company).subscribe(ret => {
      if(ret)
        this.router.navigate(['/profile']);
    });
  }
}
