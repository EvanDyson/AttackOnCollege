import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
    
export class AppComponent {
  title = 'AttackOnCollege';
  ImagePath: string;
  BImagePath: string;
  constructor() {
    //image location
    this.ImagePath = '/assets/Images/aoc.jpeg'
    this.BImagePath = '/assets/Images/aoc.jpeg'
  }
}
