package videoControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
)

func ShowAllVideos(c *gin.Context) {
	search := c.Query("search")

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var videos *[]models.Video
	videosWithPagination := utils.GeneratePagination(c, &videos)

	videoRepository := repository.NewVideoRepository(db)

	if search != "" {
		if videos, err = videoRepository.FindAllVideosByTitle(search, videosWithPagination); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if videos, err = videoRepository.FindAllVideos(videosWithPagination); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, videosWithPagination)
}
