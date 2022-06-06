package usersController

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
)

// ShowAllUsers godoc
// @Summary Show all users
// @Description Show all users
// @Tags users
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /users [get]
func ShowAllUsers(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userPagination := utils.GeneratePagination[models.User](c)
	userRepository := repository.NewUserRepository(db)
	if err = userRepository.FindAll(&userPagination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userPagination)
}
