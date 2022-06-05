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

func (ur *UserRepository) FindAll(user *[]models.User) error {
	return ur.db.Find(user).Error
}

func (ur *UserRepository) FindUserByID(user *models.User) error {
	return ur.db.First(user).Error
}

func (ur *UserRepository) UpdateUser(user *models.User, userDataToUpdate *models.User) error {
	return ur.db.Model(user).Updates(userDataToUpdate).Error
}

func (ur *UserRepository) UpdateUserPassword(userID uint, password string) error {
	return ur.db.Model(&models.User{}).Where("id = ?", userID).Update("password", password).Error
}

func (ur *UserRepository) DeleteUser(user *models.User) error {
	return ur.db.Delete(user).Error
}
