import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HostsgroupListComponent } from './hostsgroup-list.component';

describe('HostsgroupListComponent', () => {
  let component: HostsgroupListComponent;
  let fixture: ComponentFixture<HostsgroupListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HostsgroupListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HostsgroupListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
