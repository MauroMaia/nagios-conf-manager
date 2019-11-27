import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HostsListComponent } from './hosts-list.component';

describe('HostsListComponent', () => {
  let component: HostsListComponent;
  let fixture: ComponentFixture<HostsListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HostsListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HostsListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
