package videos_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
	"strconv"
)

func UpdateVideo(c *gin.Context) {
	IDVideo := c.Param("id")
	IDUint, err := strconv.ParseUint(IDVideo, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var video models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		errValidationMessagesResponse := utils.GetErrValidationMessageResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errValidationMessagesResponse})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	video.ID = uint(IDUint)
	videoRepository := repository.NewVideoRepository(db)
	if _, err := videoRepository.FindByID(video.ID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	err = videoRepository.Update(&video)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Video updated successfully"})
}
