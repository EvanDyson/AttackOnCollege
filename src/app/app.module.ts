import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { AdminPageComponent } from './admin-page/admin-page.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatStepperModule } from '@angular/material/stepper';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatListModule } from '@angular/material/list';
import { MatSelectModule } from '@angular/material/select';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatNativeDateModule } from '@angular/material/core';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import {  CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { UserProfileComponent } from './user-profile/user-profile.component';
import { MatCardModule } from '@angular/material/card';
import { EditUserComponent } from './edit-user/edit-user.component';
import { AddCourseComponent } from './add-course/add-course.component';
import { AddAssignmentComponent } from './add-assignment/add-assignment.component';
import { AchievementComponent } from './achievement/achievement.component';
import { NgxMatDatetimePickerModule, NgxMatTimepickerModule, NgxMatNativeDateModule } from '@angular-material-components/datetime-picker';
import { UniversalAppInterceptor } from './http-interceptor.service';

@NgModule({
  
  declarations: [
    AppComponent,
    LandingPageComponent,
    LoginComponent,
    RegisterComponent,
    AdminPageComponent,
    UserProfileComponent,
    EditUserComponent,
    AddCourseComponent,
    AddAssignmentComponent,
    AchievementComponent
    ],
  imports: [
      RouterModule,
      HttpClientModule,
      AppRoutingModule,
      CommonModule,
    BrowserModule,
    BrowserAnimationsModule,
    MatStepperModule,
    FormsModule,
    ReactiveFormsModule,
    MatInputModule,
    MatButtonModule,
    MatListModule,
    MatNativeDateModule,
      MatDatepickerModule,
      MatIconModule,
      MatCardModule,
      MatFormFieldModule,
      MatSelectModule,
      NgxMatTimepickerModule,
      NgxMatDatetimePickerModule,
      NgxMatNativeDateModule
  ],
  schemas:[NO_ERRORS_SCHEMA],
  
  providers: [{provide: HTTP_INTERCEPTORS, useClass: UniversalAppInterceptor,multi:true}],
    bootstrap: [AppComponent]
})
export class AppModule { }
