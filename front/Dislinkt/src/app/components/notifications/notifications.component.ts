import { HttpErrorResponse } from '@angular/common/http';
import { UserService } from './../../services/user.service';
import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/model/user';
import { NotificationService } from 'src/app/services/notification.service';
import { Router } from '@angular/router';
import { Notification } from 'src/app/model/notification';

@Component({
  selector: 'app-notifications',
  templateUrl: './notifications.component.html',
  styleUrls: ['./notifications.component.css']
})
export class NotificationsComponent implements OnInit {
  isUserLoggedIn: boolean = false;
  loggedUser: User;

  areNotificationsEmpty: boolean = true;
  notifications: Array<Notification>;

  constructor(private userService: UserService, private notificationService: NotificationService, private router: Router) { }

  ngOnInit(): void {
    this.notifications = new Array<Notification>();
    this.userService.getUserByUsername(sessionStorage.getItem("username")).subscribe(user => {
      this.loggedUser = Object.values(user)[0];
      this.isUserLoggedIn = true;
      this.notificationService.getNotificationsByReceiver(this.loggedUser.username).subscribe(nots => {
        let notifications = nots['notifications'];
        for (let not of notifications) {
          this.areNotificationsEmpty = false;
          this.notifications.push(not);
        }
      })
    }, () => {
      this.router.navigate(['/login']);
    });
  }

  readNotification(notification: Notification) {
    notification.isRead = true;
    this.notificationService.editNotification(notification).subscribe(ret => {
      
    }, (error: HttpErrorResponse) => {
      alert(error.message)
    })
  }

}
