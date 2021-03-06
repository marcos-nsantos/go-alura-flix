package videoControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"net/http"
	"strconv"
)

// ShowVideo
// @Summary Show video
// @Description Show video
// @Tags videos
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Video
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /videos/{id} [get]
func ShowVideo(c *gin.Context) {
	IDVideo := c.Param("id")
	IDUint, err := strconv.ParseUint(IDVideo, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
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

	c.JSON(http.StatusOK, video)
}
