import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  isLinear = true;
  firstFormGroup: FormGroup;

  secondFormGroup: FormGroup;

  

  constructor(private _formBuilder: FormBuilder) {}

  

  ngOnInit() {

    this.firstFormGroup = this._formBuilder.group({

      username: ['', Validators.required],

      password: ['', Validators.required]

    });


  }

  

  submit(){

      console.log(this.firstFormGroup.value);

      console.log(this.secondFormGroup.value);

  }
}
