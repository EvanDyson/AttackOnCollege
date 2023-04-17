import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent {
  
  constructor(private http: HttpClient) {};
 
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
    const element6=document.getElementById("age");
    if(element!=null){
        element.innerHTML = "Name: " + data["FirstName"] + " " + data["LastName"];
    }
    if(element1!=null){
      element1.innerHTML="Username: " + data["Username"];
    }
    if(element2!=null){
      element2.innerHTML="DOB: " + data["DOB"];
    }
    if(element3!=null){
      element3.innerHTML="Email: " + data["Email"];
    }
    if(element4!=null){
      element4.innerHTML="College: " + data["College"];
    }
    if(element5!=null){
      element5.innerHTML="Major: " + data["Major"];
    }
    if(element6!=null){
      element6.innerHTML="Age: " + data["age"];
    }
  }
}
