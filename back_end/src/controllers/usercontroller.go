// Functions that control user behavior
package controllers

import (
	"net/http"

	"CEN3031-Project/back_end/src/database"
	"CEN3031-Project/back_end/src/models"

	"github.com/gin-gonic/gin"
)

// Temporary struct that contains all information that is passed in upon user registration
type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Major     string `json:"major" binding:"required"`
	College   string `json:"college" binding:"required"`
	DOB       string `json:"dob"`
}

func RegisterUser(context *gin.Context) {
	var user models.User
	var request RegisterRequest
	// Insert info into object user and check if the information provided in the request matches fields to those of object User
	if err := context.ShouldBindJSON(&request); err != nil {
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
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

func createUser(user *models.User, request *RegisterRequest) {
	user.Username = request.Username
	user.Password = request.Password
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.DOB = request.DOB
	user.Major = request.Major
	user.College = request.College
}
