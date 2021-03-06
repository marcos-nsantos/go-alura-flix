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

func (cr CategoriaRepository) FindAll(pagination *models.Pagination[models.Categoria]) error {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := cr.db.Model(&models.Categoria{}).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	return queryBuilder.Find(&pagination.Results).Count(&pagination.TotalRows).Error
}

func (cr CategoriaRepository) FindByID(categoria *models.Categoria) error {
	return cr.db.First(categoria).Error
}

func (cr CategoriaRepository) UpdateCategoria(categoria *models.Categoria, dataToUpdate *models.Categoria) error {
	return cr.db.Model(categoria).Updates(dataToUpdate).Error
}

func (cr CategoriaRepository) DeleteCategoria(categoria *models.Categoria) error {
	return cr.db.Delete(categoria).Error
}

func (cr CategoriaRepository) VideosBelongsToCategoria(IDCategoria uint, pagination *models.Pagination[models.Video]) error {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := cr.db.Model(&models.Video{}).Where("categoria_id = ?", IDCategoria).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	return queryBuilder.Find(&pagination.Results).Count(&pagination.TotalRows).Error
}
