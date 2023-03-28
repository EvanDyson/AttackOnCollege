import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent {
  postId: string;
  constructor(private http: HttpClient) {};
  username: String;
  firstname: String;
  lastname: String;
  DOB: String;
  email: String;
  college: String;
  major: String;
  data: any;
  ngOnInit() {
    this.getInfo();
  }
  getInfo(){
    this.http.get('http://localhost:1337/users/secured/token')
    .subscribe((data: any) =>{
      this.postId = JSON.stringify(data);
      console.log(this.postId);
      
    });
  }
}
