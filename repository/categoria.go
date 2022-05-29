package repository

import (
	"github.com/marcos-nsantos/alura-flix/models"
	"gorm.io/gorm"
)

type CategoriaRepository struct {
	db *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) *CategoriaRepository {
	return &CategoriaRepository{db}
}

func (cr CategoriaRepository) CreateCategoria(categoria *models.Categoria) error {
	return cr.db.Create(categoria).Error
}

func (cr CategoriaRepository) FindAll() ([]*models.Categoria, error) {
	var categorias []*models.Categoria
	err := cr.db.Find(&categorias).Error
	return categorias, err
}
