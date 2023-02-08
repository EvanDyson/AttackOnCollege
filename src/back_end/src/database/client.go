package database

import (
	"CEN3031-Project/src/back_end/src/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var UserDB, CourseDB, AchievementDB, AssignmentDB *gorm.DB
var dbError error

func Connect(databasePath string) {
	UserDB, dbError = gorm.Open(sqlite.Open(databasePath+"/users.db"), &gorm.Config{})
	AchievementDB, dbError = gorm.Open(sqlite.Open(databasePath+"/achievements.db"), &gorm.Config{})
	AssignmentDB, dbError = gorm.Open(sqlite.Open(databasePath+"/assignments.db"), &gorm.Config{})
	CourseDB, dbError = gorm.Open(sqlite.Open(databasePath+"/courses.db"), &gorm.Config{})

	if dbError != nil {
		log.Fatal("Error!")
		panic("Error connecting to database!")
	}
	log.Println("Connected to database!")
}

func Migrate() {
	UserDB.AutoMigrate(&models.User{})
	CourseDB.AutoMigrate(&models.Course{})
	AchievementDB.AutoMigrate(&models.Achievement{})
	AssignmentDB.AutoMigrate(&models.Assignment{})
	log.Println("Database Migration Completed!")
}
