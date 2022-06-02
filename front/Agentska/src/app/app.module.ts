import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RegisterComponent } from './components/register/register.component';
import { LoginComponent } from './components/login/login.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { UsersPageComponent } from './components/users-page/users-page.component';
import { AddCommentToCompanyComponent } from './components/add-comment-to-company/add-comment-to-company.component';
import { CompanyProfileComponent } from './components/company-profile/company-profile.component';
import { AddSalaryToCompanyComponent } from './components/add-salary-to-company/add-salary-to-company.component';
import { AddExperiencesOfInterviewToCompanyComponent } from './components/add-experiences-of-interview-to-company/add-experiences-of-interview-to-company.component';
import { RegisterCompanyComponent } from './components/register-company/register-company.component';
import { EditCompanyDescriptionComponent } from './components/edit-company-description/edit-company-description.component';
import { RegisterCompanyRequestsComponent } from './components/register-company-requests/register-company-requests.component';

@NgModule({
  declarations: [
    AppComponent,
    RegisterComponent,
    LoginComponent,
    UsersPageComponent,
    AddCommentToCompanyComponent,
    CompanyProfileComponent,
    AddSalaryToCompanyComponent,
    AddExperiencesOfInterviewToCompanyComponent,
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
