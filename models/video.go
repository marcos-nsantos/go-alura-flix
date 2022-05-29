package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	URL       string `json:"url"`
}
