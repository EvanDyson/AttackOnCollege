import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  isLinear = true;
  firstFormGroup: FormGroup;

  secondFormGroup: FormGroup;

  

  constructor(private _formBuilder: FormBuilder) {}

  

  ngOnInit() {

    this.firstFormGroup = this._formBuilder.group({

      name: ['', Validators.required],

      description: ['', Validators.required]

    });

    this.secondFormGroup = this._formBuilder.group({

      amount: ['', Validators.required],

      stock: ['', Validators.required]

    });

  }

  

  submit(){

      console.log(this.firstFormGroup.value);

      console.log(this.secondFormGroup.value);

  }
}
