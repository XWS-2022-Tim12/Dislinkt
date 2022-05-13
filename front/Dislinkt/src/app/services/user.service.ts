import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { User } from '../model/user';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private loginUrl: string;
  private registerUrl: string;

  constructor(private http: HttpClient) {
    this.loginUrl = 'http://localhost:8000/user/login';
    this.registerUrl = 'http://localhost:8000/user';
   }

   public login(user: User): Observable<void> {
    return this.http.post<void>(this.loginUrl, user,{withCredentials: true});
  }
  public register(user: User): Observable<object> {
    return this.http.post<object>(this.registerUrl, user,{withCredentials: true});
  }
}
