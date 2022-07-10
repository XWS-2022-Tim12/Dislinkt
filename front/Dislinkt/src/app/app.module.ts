import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { RegisterComponent } from './components/register/register.component';
import { ProfileComponent } from './components/profile/profile.component';
import { HomePageGuestComponent } from './components/home-page-guest/home-page-guest.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './material/material.module';
import { MatNativeDateModule } from '@angular/material/core';
import { OtherUserProfileComponent } from './components/other-user-profile/other-user-profile.component';
import { FollowingRequestsComponent } from './components/following-requests/following-requests.component';
import { NewPostComponent } from './components/new-post/new-post.component';
import { LinkyModule } from 'ngx-linky';
import { ShowUsersComponent } from './components/show-users/show-users.component';
import { ShowSuggestionsForFollowingComponent } from './components/show-suggestions-for-following/show-suggestions-for-following.component';


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    ProfileComponent,
    HomePageGuestComponent,
    OtherUserProfileComponent,
    FollowingRequestsComponent,
    NewPostComponent,
    ShowUsersComponent,
    ShowSuggestionsForFollowingComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    NgbModule,
    BrowserAnimationsModule,
    MatNativeDateModule,
    MaterialModule,
    LinkyModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
