package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required,notblank"`
	Email    string `json:"email,notblank"`
	Password string `json:"password,notblank,omitempty"`
}
