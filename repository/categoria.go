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

func (cr CategoriaRepository) FindByID(id uint) (*models.Categoria, error) {
	var categoria models.Categoria
	err := cr.db.First(&categoria, id).Error
	return &categoria, err
}

func (cr CategoriaRepository) UpdateCategoria(categoria *models.Categoria) error {
	return cr.db.Save(categoria).Error
}

func (cr CategoriaRepository) DeleteCategoria(categoria *models.Categoria) error {
	return cr.db.Delete(categoria).Error
}
