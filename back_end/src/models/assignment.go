package models

import (
	"math"

	"gorm.io/gorm"
)

const XP_PTS = 100

type Assignment struct {
	gorm.Model
	Title       string `form:"assignmentName"`
	Description string `form:"description"`
	DueDate     string `form:"dueDate"`
	Type        string `form:"assignmentType"`

	NumberOfPoints int     `form:"numPts"`
	Weight         float32 `form:"gradeWeight"`

	PointsEarned float32 `form:"ptsEarned"`

	// expPts = ptsEarned / NumberOfPoints * Weight * (Set number of points for each assignment - TO BE DETERMINED)
	ExperiencePoints int  `form:"expPts"`
	IsDone           bool `form:"isDone"`
}

func (assignment *Assignment) CalculateXP() {
	assignment.ExperiencePoints = int(math.Ceil((float64(assignment.PointsEarned) / float64(assignment.NumberOfPoints)) * float64(assignment.Weight) * XP_PTS))
}
