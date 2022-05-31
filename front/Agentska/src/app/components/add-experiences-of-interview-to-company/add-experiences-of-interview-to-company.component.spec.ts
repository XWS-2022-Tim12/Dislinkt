import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddExperiencesOfInterviewToCompanyComponent } from './add-experiences-of-interview-to-company.component';

describe('AddExperiencesOfInterviewToCompanyComponent', () => {
  let component: AddExperiencesOfInterviewToCompanyComponent;
  let fixture: ComponentFixture<AddExperiencesOfInterviewToCompanyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddExperiencesOfInterviewToCompanyComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddExperiencesOfInterviewToCompanyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
