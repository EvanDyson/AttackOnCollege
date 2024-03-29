package database

import (
	"AttackOnCollege/back_end/src/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var UserDB, CourseDB, AchievementDB, AssignmentDB *gorm.DB
var dbError error

func Connect(databasePath string) {
	//Below is a local Data Source Name that can be used to access a potential remote PostgreSQL Database in future implementation
	//dsn := "host=localhost user=postgres password=CENMoment123 dbname=postgres port=1337 sslmode=disable TimeZone=EST"
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
	CreateAdmin()
}

func CreateAdmin() {
	admin := models.User{
		Email:                "",
		Username:             "AOCAdmin",
		Password:             "SuperSecretP4ssFor4dmin",
		Token:                "",
		FirstName:            "",
		LastName:             "",
		Major:                "",
		College:              "",
		DOB:                  "",
		IsAdmin:              true,
		Age:                  0,
		Level:                0,
		ExperiencePoints:     0,
		CurrentCourse:        "",
		CompletedAssignments: 0,
	}
	if err := admin.HashPassword(admin.Password); err != nil {
		return
	}
	r := UserDB.Create(&admin)
	if r.Error != nil {
		fmt.Printf(r.Error.Error())
		return
	}
}
