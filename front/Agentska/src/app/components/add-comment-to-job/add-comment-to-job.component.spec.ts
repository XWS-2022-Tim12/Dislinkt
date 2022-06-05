import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddCommentToJobComponent } from './add-comment-to-job.component';

describe('AddCommentToJobComponent', () => {
  let component: AddCommentToJobComponent;
  let fixture: ComponentFixture<AddCommentToJobComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddCommentToJobComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddCommentToJobComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
