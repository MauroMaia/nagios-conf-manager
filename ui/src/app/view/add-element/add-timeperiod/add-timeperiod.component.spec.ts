import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AddTimeperiodComponent } from './add-timeperiod.component';

describe('AddTimeperiodComponent', () => {
  let component: AddTimeperiodComponent;
  let fixture: ComponentFixture<AddTimeperiodComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AddTimeperiodComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AddTimeperiodComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
