import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface assignmentStruct {
  id: number;
  title: string;
  due_date: string;
  course_code: string;
}

const AssignmentCard: assignmentStruct[] = [];

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent {
  columns: string[] = ['title', 'due_date', 'course_code']
  dataSource: any;
  constructor(private http: HttpClient) {};
 
  ngOnInit() {
    this.getInfo();
    
  }
  getInfo(){
    this.http.get('http://localhost:1337/users/secured/token')
    .subscribe((data: any) =>{
      this.setStrings(data[0]);
      this.getAssignmentInfo(data[1], data[2])
      this.dataSource = AssignmentCard
    });
  }
  getAssignmentInfo(size: any, data: any) {
    for (let i = 0; i < size; i++) {
      AssignmentCard.push({id: data[i]["ID"], title: data[i]["Title"], due_date: data[i]["DueDate"], course_code: data[i]["Course"]})
    }
  }
  setStrings(data: any){
    const element=document.getElementById("name");
    const element1=document.getElementById("username");
    const element2=document.getElementById("dob");
    const element3=document.getElementById("email");
    const element4=document.getElementById("college");
    const element5=document.getElementById("major");
    const element6=document.getElementById("age");
    const element7=document.getElementById("currCourse");
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
      element6.innerHTML="Age: " + data["Age"];
    }
    if(element7!=null){
      element7.innerHTML="Current course: " + data["CurrentCourse"];
    }
  }
}
