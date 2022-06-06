package dto

import "github.com/marcos-nsantos/alura-flix/models"

type VideoDTO struct {
	CategoriaID uint   `json:"categoriaID" biding:"omitempty"`
	Titulo      string `json:"titulo" binding:"required,notblank"`
	Descricao   string `json:"descricao" binding:"required,notblank"`
	URL         string `json:"url" binding:"required,notblank,url"`
}

func (videoDTO VideoDTO) ToVideo() *models.Video {
	return &models.Video{
		CategoriaID: videoDTO.CategoriaID,
		Titulo:      videoDTO.Titulo,
		Descricao:   videoDTO.Descricao,
		URL:         videoDTO.URL,
	}
}
