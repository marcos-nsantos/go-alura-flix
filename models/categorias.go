package models

import (
	"gorm.io/gorm"
	"time"
)

type Categoria struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Titulo    string         `json:"titulo" binding:"required,notblank"`
	Cor       string         `json:"cor" binding:"required,notblank,iscolor"`
	Video     *Video         `json:"video,omitempty"`
}
