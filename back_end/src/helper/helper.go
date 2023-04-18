package helper

import (
	"AttackOnCollege/back_end/src/database"
	"AttackOnCollege/back_end/src/models"
	"fmt"
)

func Main() {
	CreateTestAcc()
	AddAchievementToDB("First Blood!", "Congratulations on completing your first assignment in Attack on College! Keep going!", 100)
	AddAchievementToDB("Triple Kill", "3 assignments completed! Great work.", 200)
	AddAchievementToDB("Unstoppable", "After completing 10 assignments, you are truly unstoppable!", 500)
	AddAchievementToDB("Perfect score!", "Getting a 100% on an assignment", 100)
	AddAchievementToDB("Course Killer!", "Done with the first course!", 100)

	AddAchievementToAcc("First Blood!")
	AddAchievementToAcc("Triple Kill")
	AddAchievementToAcc("Unstoppable")
	AddAchievementToAcc("Perfect score!")
	AddAchievementToAcc("Course Killer!")
}

func CreateTestAcc() {
	test := models.User{
		Email:                "AOCtest@aoc.com",
		Username:             "AOCTest",
		Password:             "P4ssw0rdF0rTest",
		Token:                "",
		FirstName:            "",
		LastName:             "",
		Major:                "",
		College:              "",
		DOB:                  "",
		IsAdmin:              false,
		Age:                  0,
		Level:                0,
		ExperiencePoints:     0,
		CurrentCourse:        "",
		CompletedAssignments: 0,
	}
	if err := test.HashPassword(test.Password); err != nil {
		return
	}
	r := database.UserDB.Create(&test)
	if r.Error != nil {
		fmt.Printf(r.Error.Error())
		return
	}
}

func AddAchievementToDB(title string, description string, xp int) {
	a := models.Achievement{
		Title:            title,
		Description:      description,
		ExperiencePoints: xp,
	}
	record := database.AchievementDB.Create(&a)
	if record.Error != nil {
		fmt.Printf(record.Error.Error())
		return
	}
}

func AddAchievementToAcc(title string) {
	var a models.Achievement
	record := database.AchievementDB.Where("title = ?", title).First(&a)
	if record.Error != nil {
		fmt.Println(record.Error.Error())
		return
	}
	var test models.User
	r := database.UserDB.Where("username = ?", "AOCTest").First(&test)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
		return
	}
	test.Achievements = append(test.Achievements, int64(a.ID))
	database.UserDB.Save(&test)
}
