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
  private getAllUrl: string;
  private editBasicUrl: string;
  private editAdvancedUrl: string;
  private editPersonalUrl: string;
  private editAllUrl: string;

  constructor(private http: HttpClient) {
    this.loginUrl = 'http://localhost:8000/user/login';
    this.registerUrl = 'http://localhost:8000/user';
    this.getAllUrl = 'http://localhost:8000/user/users';
    this.editBasicUrl = 'http://localhost:8000/user/info/basic';
    this.editAdvancedUrl = 'http://localhost:8000/user/info/advanced';
    this.editPersonalUrl = 'http://localhost:8000/user/info/personal';
    this.editAllUrl = 'http://localhost:8000/user/info/all';
   }

   public login(user: User): Observable<void> {
    return this.http.post<void>(this.loginUrl, user,{withCredentials: true});
  }
  public register(user: User): Observable<object> {
    return this.http.post<object>(this.registerUrl, user,{withCredentials: true});
  }
  public getAll(): Observable<object> {
    return this.http.get<object>(this.getAllUrl,{withCredentials: true});
  }
  public editBasic(user: User): Observable<string> {
    return this.http.put<string>(this.editBasicUrl, user,{responseType: 'text' as 'json',withCredentials: true});
  }
  public editAdvanced(user: User): Observable<string> {
    return this.http.put<string>(this.editAdvancedUrl, user,{responseType: 'text' as 'json',withCredentials: true});
  }
  public editPersonal(user: User): Observable<string> {
    return this.http.put<string>(this.editPersonalUrl, user,{responseType: 'text' as 'json',withCredentials: true});
  }
  public editAll(user: User): Observable<string> {
    return this.http.put<string>(this.editAllUrl, user,{responseType: 'text' as 'json',withCredentials: true});
  }
}