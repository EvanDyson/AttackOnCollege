package controllers

import (
	"CEN3031-Project/AttackOnCollege_v0.0.1/src/database"
	"CEN3031-Project/AttackOnCollege_v0.0.1/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(context *gin.Context) {
	var user models.User
	tokenString := context.GetHeader("Authorization")

	recordUser := database.UserDB.Where("token = ?", tokenString).First(&user)
	if recordUser.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": recordUser.Error.Error()})
		context.Abort()
		return
	}
	var request = struct {
		Title string `json:"title"`
		Code  string `json:"code"`
	}{}
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	var course models.Course
	course.Title = request.Title
	course.CourseCode = request.Code
	course.IsDone = false

	recordCourse := database.CourseDB.Create(&course)
	if recordCourse.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": recordCourse.Error.Error()})
		context.Abort()
		return
	}
	user.CurrentCourse = course.CourseCode
	database.UserDB.Save(&user)

	context.JSON(http.StatusCreated, gin.H{"courseTitle": course.Title, "courseCode": course.CourseCode})
}
