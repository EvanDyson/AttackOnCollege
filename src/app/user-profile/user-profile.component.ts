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
  username: string;
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
      this.setStrings(data);
      
    });
  }
  setStrings(data: any){
    const element=document.getElementById("name");
    const element1=document.getElementById("username");
    const element2=document.getElementById("dob");
    const element3=document.getElementById("email");
    const element4=document.getElementById("college");
    const element5=document.getElementById("major");
    if(element!=null){
      element.innerHTML="Name: " + data["Firstname"] + " " + data["LastName"];

    }
    if(element1!=null){
      element1.innerHTML="Name: " + data["Firstname"] + " " + data["LastName"];

    }
    if(element2!=null){
      element2.innerHTML="Name: " + data["Firstname"] + " " + data["LastName"];

    }
    if(element3!=null){
      element3.innerHTML="Name: " + data["Firstname"] + " " + data["LastName"];

    }
    if(element4!=null){
      element4.innerHTML="Name: " + data["Firstname"] + " " + data["LastName"];

    }
    if(element5!=null){
      element5.innerHTML="Name: " + data["Firstname"] + " " + data["LastName"];

    }

  }
}
