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
  
      var formData: any=new FormData();
      this.addData(formData);
      this.http.post('http://localhost:1337/users/register', formData)
      .subscribe(data =>{
        this.postId=JSON.stringify(data);
        console.log(this.postId);
      });
    }
    
  addData(formData: FormData){
      formData.append('firstName', this.firstFormGroup.get('firstName')?.value);
      formData.append('lastName', this.firstFormGroup.get('lastName')?.value);
      formData.append('username', this.firstFormGroup.get('username')?.value);
      formData.append('email', this.firstFormGroup.get('email')?.value);
      formData.append('password', this.firstFormGroup.get('pasword')?.value);
      formData.append('dob', this.secondFormGroup.get('dob')?.value);
      formData.append('college', this.secondFormGroup.get('college')?.value);
      formData.append('major', this.secondFormGroup.get('major')?.value);
  }
}