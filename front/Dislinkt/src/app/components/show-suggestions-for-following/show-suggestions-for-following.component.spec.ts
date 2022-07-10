import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowSuggestionsForFollowingComponent } from './show-suggestions-for-following.component';

describe('ShowSuggestionsForFollowingComponent', () => {
  let component: ShowSuggestionsForFollowingComponent;
  let fixture: ComponentFixture<ShowSuggestionsForFollowingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ShowSuggestionsForFollowingComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ShowSuggestionsForFollowingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
