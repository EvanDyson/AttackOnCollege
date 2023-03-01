import { Component, NgModule, NO_ERRORS_SCHEMA, OnInit } from '@angular/core';
import {FormBuilder,FormGroup, FormControl, Validators} from '@angular/forms';
import { Router } from '@angular/router';
import {  CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  
})
export class LoginComponent implements OnInit {
  
    hide = true;
    loginForm: FormGroup;
    postId: string;

    constructor(private http: HttpClient, private _formBuilder: FormBuilder) { }
    
  ngOnInit() {
    this.loginForm = this._formBuilder.group({
      username: new FormControl('', Validators.required),
      password: new FormControl('', Validators.required)
    })
  }

    login() {
        var formData: any = new FormData();
        this.addData(formData);
        this.http.post('http//localhost:1337/users/register', formData)
            .subscribe(data => {
                this.postId = JSON.stringify(data);
                console.log(this.postId);
            });
    }   
  
    addData(formData: FormData) {
        formData.append('username', this.loginForm.get('username')?.value);
        formData.append('password', this.loginForm.get('pasword')?.value);
    }
}
