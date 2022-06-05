package usersController

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/repository"
	"net/http"
	"strconv"
)

func ShowUser(c *gin.Context) {
	ID := c.Param("id")
	IDUint, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.FindUserByID(uint(IDUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	user.Password = ""

	c.JSON(http.StatusOK, user)
}
