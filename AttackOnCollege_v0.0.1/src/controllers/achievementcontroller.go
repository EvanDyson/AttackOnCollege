package controllers

import (
	"CEN3031-Project/AttackOnCollege_v0.0.1/src/database"
	"CEN3031-Project/AttackOnCollege_v0.0.1/src/models"
)

func GetAchievement(user *models.User, title string) {
	var achievement models.Achievement
	record := database.AchievementDB.Where("title = ?", title).First(&achievement)

	if record.Error != nil {
		//context.JSON(http.StatusNotImplemented)
		return
	}

	user.Achievements = append(user.Achievements, (int64)(achievement.ID))
}
