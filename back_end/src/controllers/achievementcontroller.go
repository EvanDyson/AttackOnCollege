package controllers

import (
	"net/http"

	"AttackOnCollege/back_end/src/database"
	"AttackOnCollege/back_end/src/models"

	"github.com/gin-gonic/gin"
)

type AchievementRequest struct {
	Title            string `form:"title"`
	Description      string `form:"description"`
	ExperiencePoints int    `form:"expPts"`
}

func GetAchievement(user *models.User, title string) {
	var achievement models.Achievement
	record := database.AchievementDB.Where("title = ?", title).First(&achievement)

	if record.Error != nil {
		//context.JSON(http.StatusNotImplemented)
		return
	}

	user.Achievements = append(user.Achievements, (int64)(achievement.ID))
}

/*** TODO: Add a function that responds to HTTP GET request for a single achievement ***/

func AddAchievement(context *gin.Context) {
	var achievement models.Achievement
	var request AchievementRequest

	if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	createAchievement(&achievement, &request)

	record := database.AchievementDB.Create(&achievement)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"title": achievement.Title, "descrip": achievement.Description, "expPts": achievement.ExperiencePoints})

}

func createAchievement(achievement *models.Achievement, request *AchievementRequest) {
	achievement.Title = request.Title
	achievement.Description = request.Description
	achievement.ExperiencePoints = request.ExperiencePoints
}

func GetAllAchievements(context *gin.Context) {
	var achievements []models.Achievement
	database.AchievementDB.Find(&achievements)
	context.IndentedJSON(http.StatusAccepted, achievements)
}

func EditAchievement(context *gin.Context) {
	var achievement models.Achievement
	var request AchievementRequest

	if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	achievement.Title = request.Title
	achievement.Description = request.Description
	achievement.ExperiencePoints = request.ExperiencePoints

	database.AchievementDB.Save(&achievement)
	context.JSON(http.StatusAccepted, achievement)
}

func DeleteAchievement(context *gin.Context) {
	var title = struct {
		Title string `form:"title"`
	}{}
	if err := context.Bind(&title); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.AchievementDB.Where("title = ?", title).Delete(&models.Achievement{})
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, "Successfully deleted achievement!")
}
