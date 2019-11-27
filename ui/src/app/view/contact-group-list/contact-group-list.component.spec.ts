import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ContactGroupListComponent } from './contact-group-list.component';

describe('ContactGroupListComponent', () => {
  let component: ContactGroupListComponent;
  let fixture: ComponentFixture<ContactGroupListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ContactGroupListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ContactGroupListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
