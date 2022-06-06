package videoControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
)

// ShowAllVideos godoc
// @Summary Show all videos
// @Description Show all videos
// @Tags videos
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Success 200 {array} models.Video
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /videos [get]
func ShowAllVideos(c *gin.Context) {
	search := c.Query("search")

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	videosPagination := utils.GeneratePagination[models.Video](c)
	videoRepository := repository.NewVideoRepository(db)

	if search != "" {
		if err = videoRepository.FindAllVideosByTitle(search, &videosPagination); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err = videoRepository.FindAllVideos(&videosPagination); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, videosPagination)
}
