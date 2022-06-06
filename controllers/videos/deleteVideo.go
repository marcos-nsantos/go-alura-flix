package videoControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"net/http"
	"strconv"
)

// DeleteVideo godoc
// @Summary Delete a video
// @Description Delete a video
// @Tags videos
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {string} string "Video deleted successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /videos/{id} [delete]
func DeleteVideo(c *gin.Context) {
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	err = videoRepository.DeleteVideo(&video)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Video deleted successfully"})
}
