import { TestBed } from '@angular/core/testing';

import { TimePeriodService } from './time-period.service';

describe('TimePeriodService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: TimePeriodService = TestBed.get(TimePeriodService);
    expect(service).toBeTruthy();
  });
});
