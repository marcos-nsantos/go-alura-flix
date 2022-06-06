package usersController

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
	"strconv"
)

// UpdateUserPassword godoc
// @Summary Update user password
// @Description Update user password
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path uint true "ID"
// @Param userData body models.UpdatePassword true "User data"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /users/{id}/password [patch]
func UpdateUserPassword(c *gin.Context) {
	userID := c.Param("id")
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var userDataPayload models.UpdatePassword
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

	userRepository := repository.NewUserRepository(db)
	if err := userRepository.UpdateUserPassword(uint(userIDUint), userDataPayload.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "password updated"})
}
