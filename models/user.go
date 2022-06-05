package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required,notblank"`
	Email    string `json:"email" binding:"required,notblank,email"`
	Password string `json:"password,omitempty" binding:"required,notblank"`
}

func (u *User) HashPassword() (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}
