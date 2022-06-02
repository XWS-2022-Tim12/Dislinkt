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
  private getAllPublicUsersUrl: string;
  private getPublicUserByUsernameUrl: string;
  private getUserByUsernameUrl: string;
  private followUrl: string;
  private acceptFollowingRequestUrl: string;
  private rejectFollowingRequestUrl: string;

  constructor(private http: HttpClient) {
    this.loginUrl = 'http://localhost:8000/user/login';
    this.registerUrl = 'http://localhost:8000/user';
    this.getAllUrl = 'http://localhost:8000/user/users';
    this.editBasicUrl = 'http://localhost:8000/user/info/basic';
    this.editAdvancedUrl = 'http://localhost:8000/user/info/advanced';
    this.editPersonalUrl = 'http://localhost:8000/user/info/personal';
    this.editAllUrl = 'http://localhost:8000/user/info/all';
    this.getAllPublicUsersUrl = 'http://localhost:8000/user/publicUsers';
    this.getPublicUserByUsernameUrl = 'http://localhost:8000/user/publicUserByUsername';
    this.getUserByUsernameUrl = 'http://localhost:8000/user/userByUsername';
    this.followUrl = 'http://localhost:8000/user/follow';
    this.acceptFollowingRequestUrl = 'http://localhost:8000/user/acceptFollowingRequest';
    this.rejectFollowingRequestUrl = 'http://localhost:8000/user/rejectFollowingRequest';
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
  public getPublicUserByUsername(username: string): Observable<object> {
    return this.http.get<object>(this.getPublicUserByUsernameUrl + '/' + username,{withCredentials: true});
  }
  public getAllPublicUsers(): Observable<object> {
    return this.http.get<object>(this.getAllPublicUsersUrl,{withCredentials: true});
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

  public getUserByUsername(username: string): Observable<object> {
    return this.http.get<object>(this.getUserByUsernameUrl + '/' + username, { withCredentials: true } );
  }

  public follow(user: User): Observable<string> {
    return this.http.put<string>(this.followUrl, user, { responseType: 'text' as 'json', withCredentials: true } );
  }

  public acceptFollowingRequest(user: User): Observable<string> {
    return this.http.put<string>(this.acceptFollowingRequestUrl, user, { responseType: 'text' as 'json', withCredentials: true } );
  }

  public rejectFollowingRequest(user: User): Observable<string> {
    return this.http.put<string>(this.rejectFollowingRequestUrl, user, { responseType: 'text' as 'json', withCredentials: true } );
  }
}
