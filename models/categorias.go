package models

import "gorm.io/gorm"

type Categoria struct {
	gorm.Model
	Titulo string `json:"titulo" binding:"required,notblank"`
	Cor    string `json:"cor" binding:"required,notblank,iscolor"`
}
