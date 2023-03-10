package models

import (
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// Login fields
	Username string `form:"username" gorm:"unique"`
	Password string `form:"password"`
	Token    string `form:"token"`

	// Registration fields
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
	Email     string `form:"email" gorm:"unique"`
	Major     string `form:"major" binding:"required"`
	College   string `form:"college" binding:"required"`
	DOB       string `form:"dob"`

	// Profile fields
	Level                int           `form:"level"`
	ExperiencePoints     int           `form:"expPts"`
	Achievements         pq.Int64Array `form:"achievements" gorm:"type:integer[]"`
	CurrentCourse        string        `form:"currCourse"`
	CompletedAssignments int
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPass))
	return err
}
