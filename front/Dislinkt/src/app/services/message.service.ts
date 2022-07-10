import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Message } from '../model/message';

@Injectable({
  providedIn: 'root'
})
export class MessageService {
  private createMessageUrl: string;
  private getMessagesUrl: string;

  constructor(private http: HttpClient) { 
    this.createMessageUrl = 'http://localhost:8000/user/message/newMessage';
    this.getMessagesUrl = 'http://localhost:8000/user/message/messages';
  }

  public createMessage(message: Message): Observable<string> {
    return this.http.post<string>(this.createMessageUrl, message,{withCredentials: true});
  }

  public getMessages(sender: string, receiver: string): Observable<Array<Message>> {
    return this.http.get<Array<Message>>(this.getMessagesUrl + '/' + sender + '/' + receiver,{withCredentials: true});
  }
}
