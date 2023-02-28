package models

import (
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// Login fields
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Token    string `json:"token"`

	// Registration fields
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique"`
	Major     string `json:"major" binding:"required"`
	College   string `json:"college" binding:"required"`
	DOB       string `json:"dob"`

	// Profile fields
	Level                int           `json:"level"`
	ExperiencePoints     int           `json:"expPts"`
	Achievements         pq.Int64Array `json:"achievements" gorm:"type:integer[]"`
	CurrentCourse        string        `json:"currCourse"`
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
