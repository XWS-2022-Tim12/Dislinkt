import { RegisterCompanyRequestsComponent } from './components/register-company-requests/register-company-requests.component';
import { EditCompanyDescriptionComponent } from './components/edit-company-description/edit-company-description.component';
import { RegisterCompanyComponent } from './components/register-company/register-company.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddCommentToJobComponent } from './components/add-comment-to-job/add-comment-to-job.component';
import { AddExperiencesOfInterviewToJobComponent } from './components/add-experiences-of-interview-to-job/add-experiences-of-interview-to-job.component';
import { AddSalaryToJobComponent } from './components/add-salary-to-job/add-salary-to-job.component';
import { JobProfileComponent } from './components/job-profile/job-profile.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { UsersPageComponent } from './components/users-page/users-page.component';
import { AddNewJobComponent } from './components/add-new-job/add-new-job.component';

const routes: Routes = [
  { path: '', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'profile', component: UsersPageComponent },
  { path: 'add-comment', component: AddCommentToJobComponent },
  { path: 'add-salary', component: AddSalaryToJobComponent },
  { path: 'add-experience', component: AddExperiencesOfInterviewToJobComponent },
  { path: 'job-profile', component: JobProfileComponent },
  { path: 'register-company', component: RegisterCompanyComponent },
  { path: 'edit-company-description/:name', component: EditCompanyDescriptionComponent },
  { path: 'company-requests', component: RegisterCompanyRequestsComponent },
  { path: 'new-job', component: AddNewJobComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
