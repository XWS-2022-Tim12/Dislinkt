import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddSalaryToJobComponent } from './add-salary-to-job.component';

describe('AddSalaryToJobComponent', () => {
  let component: AddSalaryToJobComponent;
  let fixture: ComponentFixture<AddSalaryToJobComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddSalaryToJobComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddSalaryToJobComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
