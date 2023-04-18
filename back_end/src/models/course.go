package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model

	Title      string `form:"courseName"`
	CourseCode string `form:"courseCode"`

	FinalGrade string `form:"finalGrade"`

	// expPts = Sum(assignment expPts) * FinalGrade + Set Pts for finishing a course
	ExperiencePoints int  `form:"expPts"`
	IsDone           bool `form:"isDone"`
}
