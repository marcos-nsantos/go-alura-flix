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
