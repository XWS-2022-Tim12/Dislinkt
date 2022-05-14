import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from '../model/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private loginUrl: string;
  private registerUrl: string;

  constructor(private http: HttpClient) {
    this.loginUrl = 'http://localhost:8080/user/login';
    this.registerUrl = 'http://localhost:8080/user';
   }

   public login(user: User): Observable<void> {
    return this.http.post<void>(this.loginUrl, user,{withCredentials: true});
  }
  public register(user: User): Observable<void> {
    return this.http.post<void>(this.registerUrl, user,{withCredentials: true});
  }
  
}
