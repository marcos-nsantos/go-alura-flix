package dto

import "github.com/marcos-nsantos/alura-flix/models"

func UserCreationToUser(userCreation *models.UserCreation) *models.User {
	user := &models.User{
		Name:     userCreation.Name,
		Email:    userCreation.Email,
		Password: userCreation.Password,
	}
	return user
}

func UpdateUserToUser(userUpdate *models.UpdateUser) *models.User {
	user := &models.User{
		Name:  userUpdate.Name,
		Email: userUpdate.Email,
	}
	return user
}
