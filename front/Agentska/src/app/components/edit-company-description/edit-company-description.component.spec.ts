import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditCompanyDescriptionComponent } from './edit-company-description.component';

describe('EditCompanyDescriptionComponent', () => {
  let component: EditCompanyDescriptionComponent;
  let fixture: ComponentFixture<EditCompanyDescriptionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EditCompanyDescriptionComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EditCompanyDescriptionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
