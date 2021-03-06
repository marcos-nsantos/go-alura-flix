package videoControllers

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

// UpdateVideo godoc
// @Summary Update a video
// @Description Update a video
// @Tags videos
// @Accept  json
// @Produce  json
// @Param id path string true "Video ID"
// @Param video body dto.VideoDTO true "Video data"
// @Success 200 {object} models.Video
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /videos/{id} [put]
func UpdateVideo(c *gin.Context) {
	IDVideo := c.Param("id")
	IDUint, err := strconv.ParseUint(IDVideo, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var videoToUpdate dto.VideoDTO
	if err := c.ShouldBindJSON(&videoToUpdate); err != nil {
		errValidationMessagesResponse := utils.GetErrValidationMessageResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errValidationMessagesResponse})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var video models.Video
	video.ID = uint(IDUint)

	videoRepository := repository.NewVideoRepository(db)
	if err := videoRepository.FindVideoByID(&video); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "video not found"})
		return
	}

	videoWithDataToUpdated := videoToUpdate.ToVideo()
	if err := videoRepository.UpdateVideo(&video, videoWithDataToUpdated); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, video)
}
