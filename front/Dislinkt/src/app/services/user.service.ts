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
  private blockUserUrl: string;
  private getSuggestionsUrl: string;
  private getUsernamesInInboxUrl: string;
  private changeNotificationsUrl: string;
  private changeNotificationsUsersUrl: string;
  private changeNotificationsMessagesUrl: string;

  constructor(private http: HttpClient) {
    this.loginUrl = 'http://localhost:8000/user/login';
    this.registerUrl = 'http://localhost:8000/user/register';
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
    this.blockUserUrl = 'http://localhost:8000/user/blockUser';
    this.getSuggestionsUrl = 'http://localhost:8000/user/suggestions/users';
    this.getUsernamesInInboxUrl = 'http://localhost:8000/user/message/messages';
    this.changeNotificationsUrl = 'http://localhost:8000/user/changeNotifications';
    this.changeNotificationsUsersUrl = 'http://localhost:8000/user/changeNotificationsUsers';
    this.changeNotificationsMessagesUrl = 'http://localhost:8000/user/changeNotificationsMessages';
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

  public blockUser(user: User): Observable<string> {
    return this.http.put<string>(this.blockUserUrl, user,{responseType: 'text' as 'json',withCredentials: true});
  }

  public getSuggestions(): Observable<Array<User>> {
    return this.http.get<Array<User>>(this.getSuggestionsUrl,{withCredentials: true});
  }

  public getUsernamesInInbox(username: string): Observable<Array<User>> {
    return this.http.get<Array<User>>(this.getUsernamesInInboxUrl + '/' + username,{withCredentials: true});
  }
  
  public changeNotifications(user: User): Observable<string> {
    return this.http.put<string>(this.changeNotificationsUrl, user, { responseType: 'text' as 'json', withCredentials: true } );
  }

  public changeNotificationsUsers(user: User): Observable<string> {
    return this.http.put<string>(this.changeNotificationsUsersUrl, user, { responseType: 'text' as 'json', withCredentials: true } );
  }

  public changeNotificationsMessages(user: User): Observable<string> {
    return this.http.put<string>(this.changeNotificationsMessagesUrl, user, { responseType: 'text' as 'json', withCredentials: true } );
  }
}
