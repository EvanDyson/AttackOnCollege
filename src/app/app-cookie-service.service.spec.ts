import { TestBed } from '@angular/core/testing';

import { AppCookieService } from './app-cookie-service.service';

describe('AppCookieServiceService', () => {
  let service: AppCookieService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AppCookieService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
