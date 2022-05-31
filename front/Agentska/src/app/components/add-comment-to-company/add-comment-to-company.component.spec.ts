import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddCommentToCompanyComponent } from './add-comment-to-company.component';

describe('AddCommentToCompanyComponent', () => {
  let component: AddCommentToCompanyComponent;
  let fixture: ComponentFixture<AddCommentToCompanyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddCommentToCompanyComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddCommentToCompanyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
