import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AbstractControl, FormControl, FormGroup, FormBuilder, ValidationErrors, ValidatorFn, Validators } from '@angular/forms';

@Component({
  selector: 'app-edit-user',
  templateUrl: './edit-user.component.html',
  styleUrls: ['./edit-user.component.css']
})
export class EditUserComponent {
  
  editUser: FormGroup;
    postId: string;
    firstName: string;
    lastName: string;
    username: string;
    email: string;
    dob: string;
    college: string;
    major: string;

  constructor(private http: HttpClient, private _formBuilder: FormBuilder) {}

    ngOnInit() {
    this.editUser = this._formBuilder.group({
      firstName: new FormControl(''),

      lastName: new FormControl(''),

      // will need to add in a unique username checking function
      username: new FormControl(''),
      
      email: new FormControl(''),

      dob: new FormControl(''),

      college: new FormControl(''),
      
      major: new FormControl('')
    });


  }
    // need to create a post request to update the user's information if necessary
    save() {
      var formData: any = new FormData();
      this.addData(formData);
      //need to create new post function for editting users?
      this.http.put('http://localhost:1337/users/secured/token', formData)
      .subscribe(data =>{
        this.postId = JSON.stringify(data);
        console.log(this.postId);
      });
        /*
        if (this.editUser.get('firstName')?.value == '')
            console.log("Blank first name")
        else if (this.editUser.get('firstName')?.value != '')
            console.log(this.editUser.get('firstName')?.value)
        */
      // move window back to the profile
      //window.location.pathname = './profile'
    }
    addData(formData: FormData) {
        if (this.editUser.get('firstName')?.value != '')
            formData.append('firstName', this.editUser.get('firstName')?.value);
        if (this.editUser.get('lastName')?.value != '')
            formData.append('lastName', this.editUser.get('lastName')?.value);
        if (this.editUser.get('username')?.value != '')
            formData.append('username', this.editUser.get('username')?.value);
        if (this.editUser.get('email')?.value != '')
            formData.append('email', this.editUser.get('email')?.value);
        if (this.editUser.get('dob')?.value != '')
            formData.append('dob', this.editUser.get('dob')?.value);
        if (this.editUser.get('college')?.value != '')
            formData.append('college', this.editUser.get('college')?.value);
        if (this.editUser.get('major')?.value != '')
            formData.append('major', this.editUser.get('major')?.value);
    }
}
