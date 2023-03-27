import { Injectable, Inject, Optional } from '@angular/core';
import { HttpInterceptor, HttpHandler, HttpRequest } from '@angular/common/http';
import { AppCookieService } from './app-cookie-service.service'; 
@Injectable()
export class UniversalAppInterceptor implements HttpInterceptor {

  constructor( private cookieService: AppCookieService) { }

  intercept(req: HttpRequest<any>, next: HttpHandler) {
    const token = this.cookieService.get('aocCookie');
    req = req.clone({
      url:  req.url,
      setHeaders: {
        Authorization: `Bearer ${token}`
      }
    });
    return next.handle(req);
  }
}