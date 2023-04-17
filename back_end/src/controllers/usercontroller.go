// Functions that control user behavior
package controllers

import (
	"net/http"

	"AttackOnCollege/back_end/src/database"
	"AttackOnCollege/back_end/src/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Temporary struct that contains all information that is passed in upon user registration
type RegisterRequest struct {
	Email     string `form:"email" binding:"required"`
	Username  string `form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
	Major     string `form:"major" binding:"required"`
	College   string `form:"college" binding:"required"`
	DOB       string `form:"dob"`
}

func RegisterUser(context *gin.Context) {
	var user models.User
	var request RegisterRequest
	// Insert info into object user and check if the information provided in the request matches fields to those of object User
	if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// Enter provided information into the user variable
	createUser(&user, &request)

	// Hash the password provided and check if any errors were thrown in the process
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	//	Create an instance of user in the database
	record := database.UserDB.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	// Toss back a response with the ID of the user created, along with the email and username used for the profile
	context.JSON(http.StatusCreated, gin.H{"email": user.Email, "username": user.Username})
}

// Formats given request string; every request does not require first 4 characters and any characters after 15 (specific day and timezone, respectively)
// (Function works as intended, just not called in correct position at the moment)
func formatDate(dob string) string {
	var newDOB string
	chars := []rune(dob)
	for i := 0; i < 15; i++ {
		if i > 3 {
			newDOB += string(chars[i])
		}
	}
	return newDOB
}

func extractAge(DOB string) int {
	now := time.Now()
	dob, err := time.Parse("Jan 02 2006", DOB)
	if err != nil {
		return -1
	}
	age := dob.Year() - now.Year()
	birthDay := dob.YearDay()
	currentDay := now.YearDay()
	if isLeap(dob.Year()) && !isLeap(now.Year()) && birthDay >= 60 {
		return birthDay - 1
	}
	if isLeap(now.Year()) && !isLeap(dob.Year()) && currentDay >= 60 {
		return birthDay + 1
	}
	if now.YearDay() < birthDay {
		age -= 1
	}
	return age
}

func isLeap(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func createUser(user *models.User, request *RegisterRequest) {
	user.Username = request.Username
	user.Password = request.Password
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.DOB = formatDate(request.DOB)
	user.Major = request.Major
	user.College = request.College
	user.Age = extractAge(user.DOB)
}
