import { TestBed } from '@angular/core/testing';

import { HostDataService } from './host-data.service';

describe('HostDataService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: HostDataService = TestBed.get(HostDataService);
    expect(service).toBeTruthy();
  });
});
