import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-achievement',
  templateUrl: './achievement.component.html',
  styleUrls: ['./achievement.component.css']
})
export class AchievementComponent {
  constructor(private http: HttpClient) {}

  postId: string;
  ngOnInit() {
    this.http.get('http://localhost:1337/users/achievement')
    .subscribe(data =>{
      this.postId = JSON.stringify(data);
      console.log(this.postId);
    })
  }


}
