package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Titulo    string `json:"titulo" binding:"required,notblank"`
	Descricao string `json:"descricao" binding:"required,notblank"`
	URL       string `json:"url" binding:"required,notblank,url"`
}
