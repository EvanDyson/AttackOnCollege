package models

import "gorm.io/gorm"

type Achievement struct {
	gorm.Model
	
	Title string `json:"title"`
	Description string `json:"description"`
	ExperiencePoints int `json:"expPts"`
	IsAchieved bool `json:"isAchieved"`
}