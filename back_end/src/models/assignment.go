package models

import (
	"math"

	"gorm.io/gorm"
)

const XP_PTS = 100 

type Assignment struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`

	NumberOfPoints int `json:"numPts"`
	Weight float32 `json:"gradeWeight"`

	PointsEarned float32 `json:"ptsEarned"`

	// expPts = ptsEarned / NumberOfPoints * Weight * (Set number of points for each assignment - TO BE DETERMINED)
	ExperiencePoints int `json:"expPts"`
	IsDone bool `json:"isDone"`
}

func (assignment *Assignment) CalculateXP() {
	assignment.ExperiencePoints = int(math.Ceil((float64(assignment.PointsEarned) / float64(assignment.NumberOfPoints)) * float64(assignment.Weight) * XP_PTS))
}