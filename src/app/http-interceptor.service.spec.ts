import { TestBed } from '@angular/core/testing';

import { UniversalAppInterceptor } from './http-interceptor.service';

describe('HttpInterceptorService', () => {
  let service: UniversalAppInterceptor;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UniversalAppInterceptor);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
