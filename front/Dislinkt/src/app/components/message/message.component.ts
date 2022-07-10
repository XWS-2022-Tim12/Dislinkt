import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Message } from 'src/app/model/message';
import { Notification } from 'src/app/model/notification';
import { User } from 'src/app/model/user';
import { MessageService } from 'src/app/services/message.service';
import { NotificationService } from 'src/app/services/notification.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-message',
  templateUrl: './message.component.html',
  styleUrls: ['./message.component.css']
})
export class MessageComponent implements OnInit {
  messages: Message[] = [];
  messageText: string = "";
  loggedUser: User;
  user: User;

  blockedNotifications: boolean;

  constructor(private route: ActivatedRoute, private userService: UserService, private messageService: MessageService, private notificationService: NotificationService) { }

  ngOnInit(): void {
    this.userService.getUserByUsername(sessionStorage.getItem("username")).subscribe(user => {
      this.loggedUser = Object.values(user)[0];
      this.route.params.subscribe(params => {
        this.userService.getUserByUsername(params['username']).subscribe(user => {
          this.user = Object.values(user)[0];
          this.messageService.getMessages(this.loggedUser.username, this.user.username).subscribe(messages => {
            this.messages = messages;
            for (let u of this.loggedUser.notificationOffMessages) {
              if (u == this.user.username) {
                this.blockedNotifications = true;
                break;
              }
            }
          });
        });
      });
    });
  }

  message() {
    if (this.messageText) {
      let message = new Message();
      message.text = this.messageText;
      message.senderUsername = this.loggedUser.username;
      message.receiverUsername = this.user.username;

      let exists = false;
      for (let follow of this.loggedUser.followingUsers) {
        if (follow === this.user.username) {
          exists = true;
        }
      }

      if (!exists) {
        alert("You are not following this user");
        return;
      }
      if (this.messages == null) {
        this.messages = [];
      }
      this.messages.push(message)
      
      this.messageService.createMessage(message).subscribe(ret => {
        
      })

      let notification = new Notification();
      notification.sender = this.loggedUser.username;
      notification.receiver = this.user.username;
      notification.creationDate = new Date();
      notification.notificationType = "message";
      notification.description = "User " + this.loggedUser.username + " sent message to " + this.user.username + ".";
      notification.isRead = false;
      this.notificationService.addNewNotification(notification).subscribe(ret => {

      });
    }
  }

  blockNotifications(username: string) {
    let user = new User();
    user.id = this.loggedUser.id;
    user.username = username;
    this.userService.changeNotificationsMessages(user).subscribe(ret => {
      window.location.reload();
    })
  }
}
