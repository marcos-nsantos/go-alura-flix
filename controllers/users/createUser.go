package usersController

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/dto"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
)

// CreateUser godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserCreation true "Create user"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var userDataPayload models.UserCreation

	if err := c.BindJSON(&userDataPayload); err != nil {
		errValidationMessagesResponse := utils.GetErrValidationMessageResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errValidationMessagesResponse})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userDataPayload.Password, err = models.HashUserPassword(&userDataPayload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := dto.UserCreationToUser(&userDataPayload)
	userRepository := repository.NewUserRepository(db)
	if err := userRepository.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
