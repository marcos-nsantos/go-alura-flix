package models

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	CategoriaID uint           `json:"categoriaID" gorm:"default:1"`
	Titulo      string         `json:"titulo" binding:"required,notblank"`
	Descricao   string         `json:"descricao" binding:"required,notblank"`
	URL         string         `json:"url" binding:"required,notblank,url"`
}
