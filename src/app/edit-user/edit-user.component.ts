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

  constructor(private http: HttpClient, private _formBuilder: FormBuilder) {}

  ngOnInit() {
    this.editUser = this._formBuilder.group({
      firstName: new FormControl('John'),

      lastName: new FormControl('Doe'),

      // will need to add in a unique username checking function
      username: new FormControl('JohnnyDoeBoy'),
      
      email: new FormControl('JohnDoe@email.com'),

      dob: new FormControl(''),

      college: new FormControl('University of Florida'),
      
      major: new FormControl('Computer Science')
    });


  }
    // need to create a post request to update the user's information if necessary
    save() {
      var formData: any = new FormData();
      this.addData(formData);
      //need to create new post function for editting users?
      this.http.post('http://localhost:1337/users/edit', formData)
      .subscribe(data =>{
        this.postId = JSON.stringify(data);
        console.log(this.postId);
      });

      // move window back to the profile
      window.location.pathname = './profile'
    }
    addData(formData: FormData) {
      formData.append('firstName', this.editUser.get('firstName')?.value);
      formData.append('lastName', this.editUser.get('lastName')?.value);
      formData.append('username', this.editUser.get('username')?.value);
      formData.append('email', this.editUser.get('email')?.value);
      formData.append('dob', this.editUser.get('dob')?.value);
      formData.append('college', this.editUser.get('college')?.value);
      formData.append('major', this.editUser.get('major')?.value);
    }
}
