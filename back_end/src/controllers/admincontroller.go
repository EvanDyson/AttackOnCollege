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

//Allows users to be directly edited from master list on admin account, to be displayed at a later date
func AdminEditUser(context *gin.Context) {
  tokenString := context.GetHeader("Authorization")
  var admin models.User
  var user models.User
  record := database.UserDB.Where("token = ?", tokenString).First(&admin)
  if record.Error != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
  }
  //Determines desired user to be edited through email search of database
  var email = struct {
    Email string `form:"email" gorm:"unique"`
  }{}
  var request = struct {
    Username string `form:"username" gorm:"unique"`
    FirstName string `form:"firstName"`
		LastName  string `form:"lastName"`
  }{}
  if admin.IsAdmin {
    record := database.UserDB.Where("email = ?", email.Email).First(&user)
	  if record.Error != nil {
		  context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		  context.Abort()
		  return
	  }

    if err := context.Bind(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

    user.Username = request.Username
    user.FirstName = request.FirstName
    user.LastName = request.LastName

    database.UserDB.Save(&user)
    //Password and token do not need to be hidden for admin user purposes
    context.JSON(http.StatusAccepted, user)
  }
}

//Returns all users to be displayed on an admin page implemented at later date
func AdminGetAllUsers(context *gin.Context) {
  var admin models.User
  var users []models.User

  if admin.IsAdmin {
    database.UserDB.Find(&users)
    context.IndentedJSON(http.StatusAccepted, users)
  }
}
