import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddSalaryToCompanyComponent } from './add-salary-to-company.component';

describe('AddSalaryToCompanyComponent', () => {
  let component: AddSalaryToCompanyComponent;
  let fixture: ComponentFixture<AddSalaryToCompanyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddSalaryToCompanyComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddSalaryToCompanyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
