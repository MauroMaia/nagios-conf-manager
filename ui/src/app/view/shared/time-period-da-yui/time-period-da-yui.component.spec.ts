import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TimePeriodDaYUIComponent } from './time-period-da-yui.component';

describe('TimePeriodDaYUIComponent', () => {
  let component: TimePeriodDaYUIComponent;
  let fixture: ComponentFixture<TimePeriodDaYUIComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TimePeriodDaYUIComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TimePeriodDaYUIComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
