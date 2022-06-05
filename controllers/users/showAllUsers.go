package usersController

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
)

func ShowAllUsers(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var users *[]models.User
	usersWithPagination := utils.GeneratePagination(c, &users)

	userRepository := repository.NewUserRepository(db)
	users, err = userRepository.FindAll(usersWithPagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, user := range *users {
		user.Password = ""
	}

	c.JSON(http.StatusOK, usersWithPagination)
}
