import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export interface AchievementStruct {
    title: string;
    description: string;
    XPgain: number;
}

const AchievementCard: AchievementStruct[] = [];

@Component({
  selector: 'app-achievement',
  templateUrl: './achievement.component.html',
  styleUrls: ['./achievement.component.css']
})
export class AchievementComponent {
  constructor(private http: HttpClient) {}
  displayedColumns: string[] = ['title', 'description', /*'progress',*/ 'XPgain'];
  dataSource: any;
  
  ngOnInit() {
    this.http.get('http://localhost:1337/users/secured/achievements')
    .subscribe(data =>{
      this.getAchievement(data);
      this.dataSource=AchievementCard;
    })
    }

    getAchievement(data: any){
      let size = data[0];
      for(let i = 0; i < size; i++){
        AchievementCard.push({title: data[1][i]["Title"],description: data[1][i]["Description"],XPgain: data[1][i]["ExperiencePoints"]});
      }
    }
}
