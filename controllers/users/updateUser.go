package usersController

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/dto"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
	"strconv"
)

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path uint true "ID"
// @Param user body models.UpdateUser true "User"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var userDataToUpdate models.UpdateUser
	if err := c.BindJSON(&userDataToUpdate); err != nil {
		errValidationMessagesResponse := utils.GetErrValidationMessageResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errValidationMessagesResponse})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	user.ID = uint(userIDUint)
	userRepository := repository.NewUserRepository(db)
	if err := userRepository.FindUserByID(&user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userWithDataToUpdated := dto.UpdateUserToUser(&userDataToUpdate)
	if err := userRepository.UpdateUser(&user, userWithDataToUpdated); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
