package controllers

import (
	"AttackOnCollege/back_end/src/database"
	"AttackOnCollege/back_end/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin functions

// Administrator deleting/removing user accounts
func AdminDeleteUser(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	var admin models.User
	record := database.UserDB.Where("token = ?", tokenString).First(&admin)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	var request = struct {
		Username string `form:"username"`
	}{}
	if admin.IsAdmin {
		if err := context.Bind(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		record := database.UserDB.Where("username = ?", request.Username).Delete(&models.User{})
		if record.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, "Succesully deleted user!")
	}
}
