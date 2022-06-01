import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageGuestComponent } from './components/home-page-guest/home-page-guest.component';
import { LoginComponent } from './components/login/login.component';
import { OtherUserProfileComponent } from './components/other-user-profile/other-user-profile.component';
import { ProfileComponent } from './components/profile/profile.component';
import { RegisterComponent } from './components/register/register.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'profile', component: ProfileComponent },
  { path: '', component: HomePageGuestComponent },
  { path: 'user/:username', component: OtherUserProfileComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
