import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AbstractControl, FormControl, FormGroup, FormBuilder, ValidationErrors, ValidatorFn, Validators } from '@angular/forms';
import { MatDatepicker } from '@angular/material/datepicker';
import { MatSelectionList } from '@angular/material/list';

@Component({
  selector: 'app-add-assignment',
  templateUrl: './add-assignment.component.html',
  styleUrls: ['./add-assignment.component.css']
})
export class AddAssignmentComponent {
  assignmentForm: FormGroup;
    
  postId: string;
  assignmentType: FormControl<any>;

  constructor(private http: HttpClient, private _formBuilder: FormBuilder) {}

  ngOnInit() {
    this.assignmentForm = this._formBuilder.group({
      assignmentName: new FormControl('', Validators.required),
      assignmentType: new FormControl('', Validators.required),
      dueDate: new FormControl('', Validators.required),
      desc: new FormControl('', Validators.required),
      points: new FormControl('',Validators.required),
      weight: new FormControl('',Validators.required),

    })
  }

  createAssignment() {
    var formData: any = new FormData();
    this.addData(formData);
    this.http.post('http://localhost:1337/users/secured/assignment', formData)
    .subscribe(data =>{
      this.postId = JSON.stringify(data);
      console.log(this.postId);
    })
    //window.location.pathname = './add-assignment';
    // For this example, we'll just log the data to the console.
   
      console.log('Assignment Name:', this.assignmentForm.get('assignmentName')?.value);
      console.log('Course Name: ', this.assignmentForm.get('courseName')?.value);
      console.log('Assignment Type: ', this.assignmentForm.get('assignmentType')?.value);
      console.log('Due Date: ', this.assignmentForm.get('dueDate')?.value);

      window.location.pathname = './add-assignment';
  }
  addData(formData: FormData) {
    formData.append('assignmentName', this.assignmentForm.get('assignmentName')?.value);
    formData.append('assignmentType', this.assignmentForm.get('assignmentType')?.value);
    formData.append('dueDate', this.assignmentForm.get('dueDate')?.value);
    formData.append('desc',this.assignmentForm.get('desc')?.value)
    formData.append('points',this.assignmentForm.get('points')?.value)
    formData.append('weight',this.assignmentForm.get('weight')?.value)
    }
}
