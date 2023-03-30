import { Component } from '@angular/core';
import { AppCookieService } from 'app/app-cookie-service.service';
@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css']
})
export class LandingPageComponent {
  title = 'AttackOnCollege';
  ImagePath: string;
  BImagePath: string;
  constructor(private cookieService: AppCookieService) {
    //image location
    this.ImagePath = '../assets/Images/aocblank.png'
    this.BImagePath = '../assets/Images/aocblank.png'
  }
  ngOnInit() {
    this.loginExist();
  }
  loginExist(){
    const element=document.getElementById("loginwhite");
    const element1=document.getElementById("login");
    const element2=document.getElementById("register");
    const element3=document.getElementById("profile");
    const element4=document.getElementById("addassignment");
    const element5 = document.getElementById("admin")
    const element6=document.getElementById("logout")
    if(element!=null && ( this.cookieService.get("aocCookie") != null)){
        element.style.visibility="hidden";
        element.setAttribute('disabled', 'disabled');
    }
    if(element1!=null && ( this.cookieService.get("aocCookie") != null)){
        element1.style.visibility = "hidden";
    }
    if(element2!=null && ( this.cookieService.get("aocCookie") != null)){
      element2.style.visibility="hidden";
    }
    if(element3!=null && ( this.cookieService.get("aocCookie") != null)){
      element3.style.visibility="visible";
    }
    if(element4!=null && ( this.cookieService.get("aocCookie") != null)){
      element4.style.visibility="visible";
    }
    if(element5!=null && ( this.cookieService.get("aocCookie") != null)){
      element5.style.visibility="visible";
      }
      if(element6!=null && ( this.cookieService.get("aocCookie") != null)){
      element6.style.visibility="visible";
    }
    }
    
    logOut() {
        this.cookieService.remove("aocCookie");
        window.location.pathname = '/';
    }
}
