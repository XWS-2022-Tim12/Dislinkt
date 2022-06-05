import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FollowingRequestsComponent } from './following-requests.component';

describe('FollowingRequestsComponent', () => {
  let component: FollowingRequestsComponent;
  let fixture: ComponentFixture<FollowingRequestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FollowingRequestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FollowingRequestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
