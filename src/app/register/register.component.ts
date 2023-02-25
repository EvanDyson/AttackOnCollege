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
  
    this.firstFormGroup = this._formBuilder.group({
  
      firstName: new FormControl('', Validators.required),

      lastName: new FormControl('', Validators.required),
  
      // will need to add in a unique username checking function
      username: new FormControl('', Validators.required),
      
      email: new FormControl('', [Validators.required, Validators.email]),
      
      password: new FormControl('', [Validators.required, Validators.pattern('^(?=.*[A-Z])(?=.*[0-9])(?=.*[a-z]).{8,}$')]),
    
      confirmPassword: new FormControl('', Validators.required)
  
    }
    );
  
    this.secondFormGroup = this._formBuilder.group({
      
      dob: ['', Validators.required],
  
        //add drop down menu to major and college for easy selection
        //also add 2 files for a bunch of majors and bunch of colleges for easy insertion to the drop down
      major: ['', Validators.required],
  
      college: ['', Validators.required]
    });
  
  }

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


    // export const confirmPasswordValidator: ValidatorFn = (control: AbstractControl): ValidationErrors | null =>
    // {
    //   const password = control.get('password');
    //   const confirmPassword = control.get('confirmPassword');

    //   return password && confirmPassword && password.value === confirmPassword.value ? { confirmPassword: true } : {confirmPassword: false};
    // };