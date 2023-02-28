package models

import "gorm.io/gorm"

type Achievement struct {
	gorm.Model

	Title string `form:"title"`
	Description string `form:"description"`
	ExperiencePoints int `form:"expPts"`
	IsAchieved bool `form:"isAchieved"`
}
