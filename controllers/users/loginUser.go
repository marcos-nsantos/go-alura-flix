package usersController

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/security"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
)

// LoginUser godoc
// @Summary Login user
// @Description Login user
// @Tags users
// @Accept  json
// @Produce  json
// @Param loginUser body models.LoginUser true "Login user"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /users/login [post]
func LoginUser(c *gin.Context) {
	var loginUser models.LoginUser
	if err := c.BindJSON(&loginUser); err != nil {
		errValidationMessagesResponse := utils.GetErrValidationMessageResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errValidationMessagesResponse})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.GetUserByEmail(loginUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid password or email"})
		return
	}

	if err := user.ComparePassword(loginUser.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password or email"})
		return
	}

	token, err := security.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
