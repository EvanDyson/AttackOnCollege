// All functions that concern behaviors of logged in users
package controllers

import (
	"net/http"

	"CEN3031-Project/back_end/src/database"
	"CEN3031-Project/back_end/src/models"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Return user information to be displayed on the user profile
func GetUser(context *gin.Context) {
	var user models.User
	// Get the authorization header from the request
	tokenString := context.GetHeader("Authorization")

	// Store information about the user with given token
	record := database.UserDB.Where("token = ?", tokenString).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	// Clear password stored in user not to reveal hashes of users passwords in the case of a potential attack
	user.Password = "Hidden"
	user.Token = ""
	// Send a response containing all the information about the user
	context.JSON(http.StatusOK, user)
}

// Modifies the information in the user profile
func EditUser(context *gin.Context) {
	var user models.User
	tokenString := context.GetHeader("Authorization")

	record := database.UserDB.Where("token = ?", tokenString).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	var request = struct {
		Username  string `form:"username" gorm:"unique"`
		FirstName string `form:"firstName"`
		LastName  string `form:"lastName"`
	}{}
	if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	user.Username = request.Username
	user.FirstName = request.FirstName
	user.LastName = request.LastName

	database.UserDB.Save(&user)
	user.Password = "Hidden"
	user.Token = ""
	context.JSON(http.StatusAccepted, user)
}

func ChangePassword(context *gin.Context) {
	var user models.User
	tokenString := context.GetHeader("Authorization")

	record := database.UserDB.Where("token = ?", tokenString).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	var request = struct {
		Password string `form:"password"`
	}{}
	if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(request.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	database.UserDB.Save(&user)
	context.JSON(http.StatusAccepted, "Successfully changed password!")
}

func DeleteUser(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	record := database.UserDB.Where("token = ?", tokenString).Delete(&models.User{})
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, "Succesully deleted user!")
}
