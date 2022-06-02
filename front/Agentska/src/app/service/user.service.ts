import { Company } from './../model/company';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from '../model/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private loginUrl: string;
  private registerUrl: string;
  private getLoggedUserUrl: string;
  private registerCompanyUrl: string;
  private acceptCompanyRegistrationRequestUrl: string;
  private changeCompanyDescriptionUrl: string;

  constructor(private http: HttpClient) {
    this.loginUrl = 'http://localhost:8080/user/login';
    this.registerUrl = 'http://localhost:8080/user';
    this.getLoggedUserUrl = 'http://localhost:8080/user/getLoggedUser';
    this.registerCompanyUrl = 'http://localhost:8080/user/registerCompany';
    this.acceptCompanyRegistrationRequestUrl = 'http://localhost:8080/user/acceptCompanyRegistrationRequest';
    this.changeCompanyDescriptionUrl = 'http://localhost:8080/user/changeCompanyDescription';
   }

   public login(user: User): Observable<void> {
    return this.http.post<void>(this.loginUrl, user,{withCredentials: true});
  }
  public register(user: User): Observable<void> {
    return this.http.post<void>(this.registerUrl, user,{withCredentials: true});
  }
  
  public getLoggedUser(): Observable<User> {
    let headers = new HttpHeaders();
    headers.append('Content-Type', 'application/json');
    let params = new HttpParams().set("username",sessionStorage.getItem('username'));
    return this.http.get<User>(this.getLoggedUserUrl, {headers: headers, params: params});
  }

  public registerCompany(company: Company): Observable<Boolean> {
    return this.http.post<Boolean>(this.registerCompanyUrl, company, { withCredentials: true } );
  }

  public acceptCompanyRegistrationRequest(company: Company): Observable<Boolean> {
    return this.http.put<Boolean>(this.acceptCompanyRegistrationRequestUrl, company, { withCredentials: true } );
  }

  public changeCompanyDescription(company: Company): Observable<Boolean> {
    return this.http.put<Boolean>(this.changeCompanyDescriptionUrl, company, { withCredentials: true } );
  }
}
