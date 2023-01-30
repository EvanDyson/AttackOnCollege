package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	
	Title string `json:"title"`
	CourseCode string `json:"courseCode" gorm:"unique"`

	FinalGrade string `json:"finalGrade"`
	Assignments pq.Int64Array `json:"assignments" gorm:"type:integer[]"`

	// expPts = Sum(assignment expPts) * FinalGrade + Set Pts for finishing a course
	ExperiencePoints int `json:"expPts"`
	IsDone bool `json:"isDone"`
}