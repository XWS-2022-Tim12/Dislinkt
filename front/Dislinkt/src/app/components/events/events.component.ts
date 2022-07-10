import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Notification } from 'src/app/model/notification';
import { User } from 'src/app/model/user';
import { NotificationService } from 'src/app/services/notification.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-events',
  templateUrl: './events.component.html',
  styleUrls: ['./events.component.css']
})
export class EventsComponent implements OnInit {

  isUserLoggedIn: boolean = false;
  loggedUser: User;

  areEventsEmpty: boolean = true;
  events: Array<Notification>;

  constructor(private userService: UserService, private notificationService: NotificationService, private router: Router) { }

  ngOnInit(): void {
    this.events = new Array<Notification>();
    this.userService.getUserByUsername(sessionStorage.getItem("username")).subscribe(user => {
      this.loggedUser = Object.values(user)[0];
      this.isUserLoggedIn = true;
      if (this.loggedUser.role == 'Admin') {
        this.notificationService.getAll().subscribe(nots => {
          let events = nots['notifications'];
          for (let not of events) {
            this.areEventsEmpty = false;
            this.events.push(not);
          }
        })
      } else {
        this.router.navigate(['/profile']);
      }
    }, () => {
      this.router.navigate(['/login']);
    });
  }

}
