package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model

	Title string `form:"title"`
	CourseCode string `form:"courseCode" gorm:"unique"`

	FinalGrade string `form:"finalGrade"`
	Assignments pq.Int64Array `form:"assignments" gorm:"type:integer[]"`

	// expPts = Sum(assignment expPts) * FinalGrade + Set Pts for finishing a course
	ExperiencePoints int `form:"expPts"`
	IsDone bool `form:"isDone"`
}
