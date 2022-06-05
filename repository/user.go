package repository

import (
	"github.com/marcos-nsantos/alura-flix/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) FindAll(pagination *models.Pagination) (*[]models.User, error) {
	var users *[]models.User
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := ur.db.Model(&models.User{}).Order(pagination.Sort)
	err := queryBuilder.Count(&pagination.TotalRows).Offset(offset).Limit(pagination.Limit).Find(&users).Error
	return users, err
}
