import { EditCompanyDescriptionComponent } from './components/edit-company-description/edit-company-description.component';
import { RegisterCompanyComponent } from './components/register-company/register-company.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddCommentToCompanyComponent } from './components/add-comment-to-company/add-comment-to-company.component';
import { AddExperiencesOfInterviewToCompanyComponent } from './components/add-experiences-of-interview-to-company/add-experiences-of-interview-to-company.component';
import { AddSalaryToCompanyComponent } from './components/add-salary-to-company/add-salary-to-company.component';
import { CompanyProfileComponent } from './components/company-profile/company-profile.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { UsersPageComponent } from './components/users-page/users-page.component';

const routes: Routes = [
  { path: '', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'profile', component: UsersPageComponent },
  { path: 'add-comment', component: AddCommentToCompanyComponent },
  { path: 'company-profile', component: CompanyProfileComponent },
  { path: 'add-salary', component: AddSalaryToCompanyComponent },
  { path: 'add-experience', component: AddExperiencesOfInterviewToCompanyComponent },
  { path: 'register-company', component: RegisterCompanyComponent },
  { path: 'edit-company-description/:name', component: EditCompanyDescriptionComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
