package database

import (
	"CEN3031-Project/back_end/src/models"
	"log"

	"gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var UserDB, CourseDB, AchievementDB, AssignmentDB *gorm.DB
var dbError error

func Connect(databasePath string) {
  dsn := "host=localhost user=postgres password=CENMoment123 dbname=postgres port=1337 sslmode=disable TimeZone=EST"
	UserDB, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	AchievementDB, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	AssignmentDB, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	CourseDB, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
