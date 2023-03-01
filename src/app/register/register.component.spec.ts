import { NO_ERRORS_SCHEMA } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RegisterComponent } from './register.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { FormControl, FormGroup } from '@angular/forms';

describe('RegisterComponent', () => {
  let component: RegisterComponent;
  let fixture: ComponentFixture<RegisterComponent>;
  let httpController:HttpTestingController;
  
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisterComponent ],
      schemas:[NO_ERRORS_SCHEMA],
      imports:[HttpClientTestingModule],
      
    })
    .compileComponents();
   
    

    fixture = TestBed.createComponent(RegisterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create "Register Page"', async() => {
    expect(component).toBeTruthy();
  });
  it('should return correct data upon registration', async()=>{
    
    component.firstFormGroup.controls['firstName'] = new FormControl('name');
    component.firstFormGroup.controls['lastName'] = new FormControl('last');
    component.firstFormGroup.controls['username'] = new FormControl('namelast');
    component.firstFormGroup.controls['email'] = new FormControl('email@ufl.edu');
    component.firstFormGroup.controls['password'] = new FormControl('12345678Test');
    component.secondFormGroup.controls['dob'] = new FormControl('01/11/2011');
    component.secondFormGroup.controls['major'] = new FormControl('CompSci');
    component.secondFormGroup.controls['college'] = new FormControl('Univeristy of Florida');
   
    component.submit();
    if(component.postId==undefined){
      component.postId="1";
     }
    console.log(component.firstFormGroup.get('firstName')?.value);
expect(component.postId).toEqual( '"email":"email@ufl.edu","username":"namelast"');

    });
  });
  
  

