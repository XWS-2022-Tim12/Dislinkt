import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterCompanyRequestsComponent } from './register-company-requests.component';

describe('RegisterCompanyRequestsComponent', () => {
  let component: RegisterCompanyRequestsComponent;
  let fixture: ComponentFixture<RegisterCompanyRequestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisterCompanyRequestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RegisterCompanyRequestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
