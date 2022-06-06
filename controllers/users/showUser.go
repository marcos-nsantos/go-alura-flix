package usersController

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"net/http"
	"strconv"
)

// ShowUser godoc
// @Summary Show user
// @Description Show user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path uint true "ID"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /users/{id} [get]
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

	var user models.User
	user.ID = uint(IDUint)

	userRepository := repository.NewUserRepository(db)
	if err := userRepository.FindUserByID(&user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	user.Password = ""

	c.JSON(http.StatusOK, user)
}
