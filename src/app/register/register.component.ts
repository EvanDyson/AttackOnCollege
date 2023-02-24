import { Component } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, FormBuilder, ValidationErrors, ValidatorFn, Validators } from '@angular/forms';
import { MatDatepicker } from '@angular/material/datepicker';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
    isLinear = true;
    hide = true;
  firstFormGroup: FormGroup;
  
  secondFormGroup: FormGroup;
  
  constructor(private _formBuilder: FormBuilder) {}
  
  ngOnInit() {
  
      // add pop ups for all error fields that says what the errors is
    this.firstFormGroup = this._formBuilder.group({
  
      firstName: ['', Validators.required],

      lastName: ['', Validators.required],
  
      username: ['', Validators.required],
      
      // add make proper validator message pop up
      email: ['', [Validators.required, Validators.email]],
      
      password: ['', [Validators.required, Validators.pattern('^(?=.*[A-Z])(?=.*[0-9])(?=.*[a-z]).{8,}$')]],
    
      // add function to make confirm must match password
      confirmPassword: ['', Validators.required]
  
    });
  
    this.secondFormGroup = this._formBuilder.group({
      
      datepicker: ['', Validators.required],
  
        //add drop down menu to major and college for easy selection
        //also add 2 files for a bunch of majors and bunch of colleges for easy insertion to the drop down
      major: ['', Validators.required],
  
      college: ['', Validators.required]
    });
  
  }

  /* NEED TO ADD ERROR MESSAGES TO THE FIELDS
  getErrorMessage() {
    if (this.firstFormGroup.hasError('required')) {
      return 'You must enter a value';
    }
    if (this.secondFormGroup.hasError('required')) {
      return 'You must enter a value';
    }
    return this.secondFormGroup.hasError('email') ? 'Not a valid email' : '';
  }
  */

  submit(){
  
    console.log(this.firstFormGroup.value);
      
      console.log(this.secondFormGroup.value);
      
      //this.http.post<any>();
      /*
        POST https://{{host}}/users/token HTTP/1.1
        content-type: application/json
        {
            "email": "b.gator@ufl.edu",
            "password": "IDon'tKnowHonestly"
        }
      */
  }
}