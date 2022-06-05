import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RegisterComponent } from './components/register/register.component';
import { LoginComponent } from './components/login/login.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { UsersPageComponent } from './components/users-page/users-page.component';
import { JobProfileComponent } from './components/job-profile/job-profile.component';
import { AddCommentToJobComponent } from './components/add-comment-to-job/add-comment-to-job.component';
import { AddExperiencesOfInterviewToJobComponent } from './components/add-experiences-of-interview-to-job/add-experiences-of-interview-to-job.component';
import { AddSalaryToJobComponent } from './components/add-salary-to-job/add-salary-to-job.component';

import { RegisterCompanyComponent } from './components/register-company/register-company.component';
import { EditCompanyDescriptionComponent } from './components/edit-company-description/edit-company-description.component';
import { RegisterCompanyRequestsComponent } from './components/register-company-requests/register-company-requests.component';


@NgModule({
  declarations: [
    AppComponent,
    RegisterComponent,
    LoginComponent,
    UsersPageComponent,
    JobProfileComponent,
    AddCommentToJobComponent,
    AddExperiencesOfInterviewToJobComponent,
    AddSalaryToJobComponent,
    RegisterCompanyComponent,
    EditCompanyDescriptionComponent,
    RegisterCompanyRequestsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
