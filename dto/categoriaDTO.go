package dto

import "github.com/marcos-nsantos/alura-flix/models"

type CategoriaDTO struct {
	Titulo string `json:"titulo" binding:"required,notblank"`
	Cor    string `json:"cor" binding:"required,notblank,iscolor"`
}

func (categoriaDTO CategoriaDTO) ToCategoria() *models.Categoria {
	return &models.Categoria{
		Titulo: categoriaDTO.Titulo,
		Cor:    categoriaDTO.Cor,
	}
}
