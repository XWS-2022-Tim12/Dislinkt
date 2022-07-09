import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Notification } from '../model/notification';

@Injectable({
  providedIn: 'root'
})
export class NotificationService {
  private addNewNotificationUrl: string;
  private editNotificationUrl: string;
  private getAllUrl: string;
  private getNotificationsBySenderUrl: string;
  private getNotificationsByReceiverUrl: string;
  private getNotificationsByTypeUrl: string;

  constructor(private http: HttpClient) {
    this.addNewNotificationUrl = 'http://localhost:8000/user/notification';
    this.editNotificationUrl = 'http://localhost:8000/user/notification/editNotification';
    this.getAllUrl = 'http://localhost:8000/notification/notifications';
    this.getNotificationsBySenderUrl = 'http://localhost:8000/notification/searchBySender';
    this.getNotificationsByReceiverUrl = 'http://localhost:8000/notification/searchByReceiver';
    this.getNotificationsByTypeUrl = 'http://localhost:8000/notification/searchByNotificationType';
   }

  public addNewNotification(notification: Notification): Observable<string> {
    return this.http.post<string>(this.addNewNotificationUrl, notification, { responseType: 'text' as 'json', withCredentials: true });
  }

  public editNotification(notification: Notification): Observable<string> {
    return this.http.put<string>(this.editNotificationUrl, notification, { responseType: 'text' as 'json', withCredentials: true });
  }
  
  public getAll(): Observable<object> {
    return this.http.get<object>(this.getAllUrl, { withCredentials: true });
  }

  public getNotificationsBySender(sender: string): Observable<Array<Notification>> {
    return this.http.get<Array<Notification>>(this.getNotificationsBySenderUrl + '/' + sender, { withCredentials: true });
  }
  
  public getNotificationsByReceiver(receiver: string): Observable<Array<Notification>> {
    return this.http.get<Array<Notification>>(this.getNotificationsByReceiverUrl + '/' + receiver, { withCredentials: true });
  }

  public getNotificationsByType(type: string): Observable<Array<object>> {
    return this.http.get<Array<object>>(this.getNotificationsByTypeUrl + '/' + type, { withCredentials: true });
  }
}
