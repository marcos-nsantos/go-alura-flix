package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Titulo    string `json:"titulo" binding:"required"`
	Descricao string `json:"descricao" binding:"required"`
	URL       string `json:"url" binding:"required,url"`
}
