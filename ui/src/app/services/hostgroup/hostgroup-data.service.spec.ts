import { TestBed } from '@angular/core/testing';

import { HostgroupDataService } from './hostgroup-data.service';

describe('HostgroupDataService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: HostgroupDataService = TestBed.get(HostgroupDataService);
    expect(service).toBeTruthy();
  });
});
