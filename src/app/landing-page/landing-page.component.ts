import { Component } from '@angular/core';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css']
})
export class LandingPageComponent {
  title = 'AttackOnCollege';
  ImagePath: string;
  BImagePath: string;
  constructor() {
    //image location
    this.ImagePath = '../assets/Images/aoc.jpeg'
    this.BImagePath = '../assets/Images/aoc.jpeg'
  }
}
