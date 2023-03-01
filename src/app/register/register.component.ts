import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
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
    
    postId: string;
  
  constructor(private http: HttpClient, private _formBuilder: FormBuilder) {}
  
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
    //, { validators: confirmPasswordValidator }
    );
  
    this.secondFormGroup = this._formBuilder.group({
      
      dob: ['', Validators.required],
  
        //add drop down menu to major and college for easy selection
        //also add 2 files for a bunch of majors and bunch of colleges for easy insertion to the drop down
      college: ['', Validators.required],
      
      major: ['', Validators.required]
    });
  
  }

  submit(){
  
      this.http.post<any>('https://{{host}}/users/token', this.firstFormGroup).subscribe(data => {
        this.postId = data.id;
    });
      this.http.post<any>('https://{{host}}/users/token', this.secondFormGroup).subscribe(data => {
        this.postId = data.id;
    });
      
    //console.log(this.firstFormGroup.value);
    //console.log(this.secondFormGroup.value);
      
      
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