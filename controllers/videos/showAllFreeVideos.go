package videoControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"net/http"
)

// ShowAllFreeVideos godoc
// @Summary Show all free videos
// @Description Show all free videos
// @Tags videos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Video
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /videos/free [get]
func ShowAllFreeVideos(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var videos []models.Video

	videoRepository := repository.NewVideoRepository(db)
	if err = videoRepository.FirstThreeVideos(&videos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)
}
