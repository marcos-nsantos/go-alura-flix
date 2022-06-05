package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Password interface {
	HashPassword() (string, error)
}

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
}

type UserCreation struct {
	Name     string `json:"name" binding:"required,notblank"`
	Email    string `json:"email" binding:"required,notblank,email"`
	Password string `json:"password" binding:"required,notblank"`
}

type UpdateUser struct {
	Name  string `json:"name" binding:"required,notblank"`
	Email string `json:"email" binding:"required,notblank,email"`
}

type UpdatePassword struct {
	Password string `json:"password" binding:"required,notblank"`
}

func (uc *UserCreation) HashPassword() (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(uc.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

func (up *UpdatePassword) HashPassword() (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(up.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

func HashUserPassword(password Password) (string, error) {
	return password.HashPassword()
}
