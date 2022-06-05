import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddExperiencesOfInterviewToJobComponent } from './add-experiences-of-interview-to-job.component';

describe('AddExperiencesOfInterviewToJobComponent', () => {
  let component: AddExperiencesOfInterviewToJobComponent;
  let fixture: ComponentFixture<AddExperiencesOfInterviewToJobComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddExperiencesOfInterviewToJobComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddExperiencesOfInterviewToJobComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
