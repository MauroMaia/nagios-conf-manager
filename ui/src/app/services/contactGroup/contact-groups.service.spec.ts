import { TestBed } from '@angular/core/testing';

import { ContactGroupsService } from './contact-groups.service';

describe('ContactGroupsService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: ContactGroupsService = TestBed.get(ContactGroupsService);
    expect(service).toBeTruthy();
  });
});
