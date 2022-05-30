import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomePageGuestComponent } from './home-page-guest.component';

describe('HomePageGuestComponent', () => {
  let component: HomePageGuestComponent;
  let fixture: ComponentFixture<HomePageGuestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HomePageGuestComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HomePageGuestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
