import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDatepicker } from '@angular/material/datepicker';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  isLinear = true;
  firstFormGroup: FormGroup;
  
  secondFormGroup: FormGroup;
  
  constructor(private _formBuilder: FormBuilder) {}
  
  ngOnInit() {
  
    this.firstFormGroup = this._formBuilder.group({
  
      firstName: ['', Validators.required],

      lastName: ['', Validators.required],
  
      college: ['', Validators.required]
  
    });
  
    this.secondFormGroup = this._formBuilder.group({
      
      datepicker: ['', Validators.required],
  
      major: ['', Validators.required],
  
      email: ['', [Validators.required, Validators.email]]
  
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
    //return this.secondFormGroup.hasError('email') ? 'Not a valid email' : '';
  }
  */

  submit(){
  
    console.log(this.firstFormGroup.value);
  
    console.log(this.secondFormGroup.value);
  }
}
