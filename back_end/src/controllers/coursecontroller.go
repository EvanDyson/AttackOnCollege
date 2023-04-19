package controllers

import (
	"AttackOnCollege/back_end/src/database"
	"AttackOnCollege/back_end/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(context *gin.Context) {
	var user models.User
	tokenString := context.GetHeader("Authorization")
	tokenString = tokenString[1 : len(tokenString)-1]
	recordUser := database.UserDB.Where("token = ?", tokenString).First(&user)
	if recordUser.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": recordUser.Error.Error()})
		context.Abort()
		return
	}
	var request = struct {
		Title string `form:"courseName"`
		Code  string `form:"courseCode"`
	}{}
	if err := context.Bind(&request); err != nil {
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
	user.CourseID = course.ID
	database.UserDB.Save(&user)

	context.JSON(http.StatusCreated, gin.H{"courseName": course.Title, "courseCode": course.CourseCode})
}

func EditCourse(context *gin.Context) {
  var user models.User
  var course models.Course
  tokenString := context.GetHeader("Authorization")
	tokenString = tokenString[1 : len(tokenString)-1]
	recordUser := database.UserDB.Where("token = ?", tokenString).First(&user)
	if recordUser.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": recordUser.Error.Error()})
		context.Abort()
		return
	}

  var request = struct {
		Title            string `form:"title"`
		CourseCode       string `form:"code"`

	}{}
  if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

  course.Title = request.Title
  course.CourseCode = request.CourseCode

  changeCourse := database.CourseDB.Save(&course)
  if changeCourse.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": changeCourse.Error.Error()})
		context.Abort()
		return
	}

  context.JSON(http.StatusAccepted, "Course Edited Successfully!")
}

func CompleteCourse(context *gin.Context) {
  var user models.User
  var course models.Course
	tokenString := context.GetHeader("Authorization")
	tokenString = tokenString[1 : len(tokenString)-1]
	recordUser := database.UserDB.Where("token = ?", tokenString).First(&user)
	if recordUser.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": recordUser.Error.Error()})
		context.Abort()
		return
	}

  var request = struct {
    CourseID   int     `form:"course"`
    FinalGrade string  `form:"finalGrade"`
  }{}
  recordCourse := database.CourseDB.Where("ID = ?", request.CourseID).First(&course)
	if recordCourse.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": recordCourse.Error.Error()})
		context.Abort()
		return
	}
  if course.IsDone {
		context.JSON(http.StatusBadRequest, "Course already completed!")
		context.Abort()
		return
	}
  course.IsDone = true
  database.CourseDB.Save(&course)
  user.ExperiencePoints += course.ExperiencePoints

  //Add user profile field to track courses completed?
  //user.CompletedCourses++
  //if user.CompletedCourses == 1 {
  //  GetAchievement(&user, "Course Killer!")
  //}

  database.UserDB.Save(&user)
  context.JSON(http.StatusAccepted, "Course completed! An awesome feat of strength!")
}
