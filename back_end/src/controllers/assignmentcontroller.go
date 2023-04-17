package controllers

import (
	"AttackOnCollege/back_end/src/database"
	"AttackOnCollege/back_end/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAssignment(c *gin.Context) {
	var user models.User
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[1 : len(tokenString)-1]
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
		Title          string  `form:"assignmentName"`
		Description    string  `form:"description"`
		DueDate        string  `form:"dueDate"`
		Type           string  `form:"assignmentType"`
		NumberOfPoints int     `form:"numPts"`
		Weight         float32 `form:"gradeWeight"`
	}{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	var assignment models.Assignment
	assignment.Title = request.Title
	assignment.Description = request.Description
	assignment.DueDate = formatDate(request.DueDate)
	assignment.Type = request.Type
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
		AssignmentID int     `form:"assignment"`
		PointsEarned float32 `form:"ptsEarned"`
	}{}
	if err := c.Bind(&request); err != nil {
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
	database.AssignmentDB.Save(&assignment)
	user.ExperiencePoints += assignment.ExperiencePoints
	user.CompletedAssignments++
	if user.CompletedAssignments == 1 {
		GetAchievement(&user, "First Blood!")
	}
	database.UserDB.Save(&user)
	c.JSON(http.StatusAccepted, "Assignment completed! Congratulations!")
	// Add function that checks if this is first completed assignment and send an achievement for it

}

func EditAssignment(c *gin.Context) {
	var user models.User
	var assignment models.Assignment
	var course models.Course
	tokenString := c.GetHeader("Authorization")

	recordUser := database.UserDB.Where("token = ?", tokenString).First(&user)
	if recordUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": recordUser.Error.Error()})
		c.Abort()
		return
	}
	recordCourse := database.CourseDB.Where("course_code = ?", user.CurrentCourse).First(&course)
	if recordCourse.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": recordCourse.Error.Error()})
		c.Abort()
		return
	}

	var request = struct {
		Title            string `form:"title"`
		Description      string `form:"description"`
		ExperiencePoints int    `form:"expPts"`
	}{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	recordAssignment := database.AssignmentDB.Where("title = ?", assignment.Title).First(&assignment)
	if recordAssignment.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": recordAssignment.Error.Error()})
		c.Abort()
		return
	}

	//Users should only be able to edit an assignment title / description / exp, not if it is completed or not
	assignment.Title = request.Title
	assignment.Description = request.Description
	assignment.ExperiencePoints = request.ExperiencePoints

	changeAssignment := database.AssignmentDB.Save(&assignment)
	if changeAssignment.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": changeAssignment.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusAccepted, "Assignment Edited Successfully!")
}
