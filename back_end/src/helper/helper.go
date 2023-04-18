package helper

import (
	"AttackOnCollege/back_end/src/database"
	"AttackOnCollege/back_end/src/models"
	"fmt"
)

func Main() {
	database.Connect("./database")
	database.Migrate()
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
	var admin models.User
	r := database.UserDB.Where("username = ?", "AOCAdmin").First(&admin)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
		return
	}
	admin.Achievements = append(admin.Achievements, int64(a.ID))
	database.UserDB.Save(&admin)
}
