package videoControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/dto"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
)

// jwt

// CreateVideo godoc
// @Summary Create a video
// @Description Create a video
// @Tags videos
// @Accept  json
// @Produce  json
// @Param video body dto.VideoDTO true "Video"
// @Success 201 {object} dto.VideoDTO
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /videos [post]
func CreateVideo(c *gin.Context) {
	var videoToCreate dto.VideoDTO

	if err := c.BindJSON(&videoToCreate); err != nil {
		errValidationMessagesResponse := utils.GetErrValidationMessageResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errValidationMessagesResponse})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	video := videoToCreate.ToVideo()
	videoRepository := repository.NewVideoRepository(db)
	if err := videoRepository.CreateVideo(video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, video)
}
