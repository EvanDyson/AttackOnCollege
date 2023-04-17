import { Component } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, FormBuilder, ValidationErrors, ValidatorFn, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-add-course',
  templateUrl: './add-course.component.html',
  styleUrls: ['./add-course.component.css']
})
export class AddCourseComponent {
    courseForm: FormGroup;
    postId: string;

    constructor(private http: HttpClient, private _formBuilder: FormBuilder) {}

    ngOnInit() {
        this.courseForm = this._formBuilder.group({
            courseCode: new FormControl('', Validators.required),
            courseName: new FormControl('', Validators.required),
            profName: new FormControl('', Validators.required)
        })
    }
    addCourse() {
        var formData: any = new FormData();
        this.addData(formData);
        this.http.post('http://localhost:1337/users/secured/course', formData)
        .subscribe(data =>{
          this.postId = JSON.stringify(data);
          console.log(this.postId);
        })
        //window.location.pathname = './add-assignment';
        // For this example, we'll just log the data to the console.
       
          console.log('Course Code: ', this.courseForm.get('courseCode')?.value);
          console.log('Course Name: ', this.courseForm.get('courseName')?.value);
          console.log('Professor Name: ', this.courseForm.get('profName')?.value);
    
          window.location.pathname = './add-course';
      }
      addData(formData: FormData) {
        formData.append('courseCode', this.courseForm.get('courseCode')?.value);
        formData.append('courseName', this.courseForm.get('courseName')?.value);
        formData.append('profName', this.courseForm.get('profName')?.value);
      }
}
