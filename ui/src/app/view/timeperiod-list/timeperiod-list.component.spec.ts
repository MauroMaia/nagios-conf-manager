import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TimeperiodListComponent } from './timeperiod-list.component';

describe('TimeperiodListComponent', () => {
  let component: TimeperiodListComponent;
  let fixture: ComponentFixture<TimeperiodListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TimeperiodListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TimeperiodListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
