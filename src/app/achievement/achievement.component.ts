import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface AchievementStruct {
    title: string;
    description: string;
    //progress: number;
    XPgain: number;
}

const AchievementCard: AchievementStruct[] = [
    { title: 'First Assignment!', description: 'Finish your first assignment.', /*progress: 0,*/ XPgain: 100 },
    { title: 'Second Assignment!', description: 'Finish your first book.', /*progress: 50,*/ XPgain: 50 },

];

@Component({
  selector: 'app-achievement',
  templateUrl: './achievement.component.html',
  styleUrls: ['./achievement.component.css']
})
export class AchievementComponent {
  constructor(private http: HttpClient) {}

  postId: string;
    
/*
  ngOnInit() {
    this.http.get('http://localhost:1337/users/achievement')
    .subscribe(data =>{
      this.postId = JSON.stringify(data);
      console.log(this.postId);
    })
    }
*/
    displayedColumns: string[] = ['title', 'description', /*'progress',*/ 'XPgain'];
    dataSource = AchievementCard;
}
