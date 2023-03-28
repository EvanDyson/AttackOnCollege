import { Component, NgModule, NO_ERRORS_SCHEMA, OnInit } from '@angular/core';
import {FormBuilder,FormGroup, FormControl, Validators} from '@angular/forms';
import { Router } from '@angular/router';
import {  CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AppCookieService } from 'app/app-cookie-service.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  
})
export class LoginComponent implements OnInit {
  
    hide = true;
    loginForm: FormGroup;
    postId: string;

    constructor(private http: HttpClient, private _formBuilder: FormBuilder,private cookieService:AppCookieService) { }
    
  ngOnInit() {
    this.loginForm = this._formBuilder.group({
      username: new FormControl('', Validators.required),
      password: new FormControl('', [Validators.required, Validators.pattern('^(?=.*[A-Z])(?=.*[0-9])(?=.*[a-z]).{8,}$')])
    })
    this.loginExist();
  }

    login() {
        var formData: any = new FormData();
        this.addData(formData);
        this.http.post('http://localhost:1337/users/token', formData)
            .subscribe(data => {
                this.postId = JSON.stringify(data);
                this.cookieService.set('aocCookie',this.postId)
                if(this.cookieService.get('aocCookie')!=null){
                  window.location.pathname = '/';
                }
            });
        
        
        
    }   
  
    addData(formData: FormData) {
        formData.append('username', this.loginForm.get('username')?.value);
        formData.append('password', this.loginForm.get('pasword')?.value);
    }
    loginExist(){
      const element=document.getElementById("loginwhite");
      const element1=document.getElementById("login");
      const element2=document.getElementById("register");
      const element3=document.getElementById("profile");
      const element4=document.getElementById("addassignment");
      const element5=document.getElementById("admin")
      if(element!=null && ( this.cookieService.get("aocCookie") != null)){
      element.style.visibility="hidden";
  
      }
      if(element1!=null && ( this.cookieService.get("aocCookie") != null)){
        element1.style.visibility="hidden";
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
    }
}
