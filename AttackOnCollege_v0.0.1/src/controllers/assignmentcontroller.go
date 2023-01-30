package controllers

import (
	"CEN3031-Project/AttackOnCollege_v0.0.1/src/database"
	"CEN3031-Project/AttackOnCollege_v0.0.1/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAssignment(c *gin.Context) {
	var user models.User
	tokenString := c.GetHeader("Authorization")
	recordUser := database.UserDB.Where("token = ?", tokenString).First(&user)
	if recordUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": recordUser.Error.Error()})
		c.Abort()
		return
	}
	var course models.Course
	recordCourse := database.CourseDB.Where("course_code = ?", user.CurrentCourse).First(&course)
	if recordUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": recordCourse.Error.Error()})
		c.Abort()
		return
	}

	var request = struct {
		Title          string  `json:"title"`
		Description    string  `json:"description"`
		NumberOfPoints int     `json:"numPts"`
		Weight         float32 `json:"gradeWeight"`
	}{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	var assignment models.Assignment
	assignment.Title = request.Title
	assignment.Description = request.Description
	assignment.NumberOfPoints = request.NumberOfPoints
	assignment.Weight = request.Weight
	assignment.IsDone = false

	recordAssignment := database.AssignmentDB.Create(&assignment)
	if recordAssignment.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": recordAssignment.Error.Error()})
		c.Abort()
		return
	}

	course.Assignments = append(course.Assignments, int64(assignment.ID))
	database.CourseDB.Save(&course)

	c.JSON(http.StatusCreated, "Assignment added!")
}

func CompleteAssignment(c *gin.Context) {
	var user models.User
	tokenString := c.GetHeader("Authorization")
	recordUser := database.UserDB.Where("token = ?", tokenString).First(&user)
	if recordUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": recordUser.Error.Error()})
		c.Abort()
		return
	}
	var course models.Course
	recordCourse := database.CourseDB.Where("course_code = ?", user.CurrentCourse).First(&course)
	if recordUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": recordCourse.Error.Error()})
		c.Abort()
		return
	}

	var request = struct {
		AssignmentID int `json:"assignment"`
		PointsEarned float32 `json:"ptsEarned"`
	}{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	var assignment models.Assignment
	recordAssignment := database.AssignmentDB.Where("ID = ?", request.AssignmentID).First(&assignment)
	if recordAssignment.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": recordAssignment.Error.Error()})
		c.Abort()
		return
	}
	if assignment.IsDone {
		c.JSON(http.StatusBadRequest, "Course already completed!")
		c.Abort()
		return
	}
	assignment.IsDone = true
	assignment.CalculateXP()
	database.AchievementDB.Save(&assignment)
	user.ExperiencePoints += assignment.ExperiencePoints
	database.UserDB.Save(&user)
	c.JSON(http.StatusAccepted, "Assignment completed! Congratulations!")
	// Add function that checks if this is first completed assignment and send an achievement for it
}