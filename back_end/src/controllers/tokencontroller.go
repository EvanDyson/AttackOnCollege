// Functions that control token behavior
package controllers

import (
	"AttackOnCollege/back_end/src/auth"
	"AttackOnCollege/back_end/src/database"
	"AttackOnCollege/back_end/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateToken is a function that deals with users logging in
// It takes in an email and a password, and generates a JWT that is used to access pages only available to page users
func GenerateToken(context *gin.Context) {
	var request = struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}{}

	var user models.User

	// Enter information given in the request into var request and check if any errors are raised in the process
	if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// Search through database to find an entry with the requested email and save the information within the var user
	record := database.UserDB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	// Check password provided against the one in the database
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	// Generate a JWT that will be used to keep the user logged in
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	user.Token = tokenString
	database.UserDB.Save(&user)
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetUsers(context *gin.Context) {
	var users []models.User
	database.UserDB.Find(&users)
	// Send a response containing all users in the database
	// No password hashes are actually sent
	for i := 0; i < len(users); i++ {
		users[i].Password = "Hidden"
		users[i].Token = ""
	}
	context.IndentedJSON(http.StatusAccepted, users)
}
