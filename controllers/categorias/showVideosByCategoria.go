package categoriaControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
	"strconv"
)

func ShowVideosByCategoria(c *gin.Context) {
	IDCategoria := c.Param("id")
	IDUint, err := strconv.ParseUint(IDCategoria, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID Inválido"})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var videos *[]models.Video
	videosWithPagination := utils.GeneratePagination(c, &videos)

	categoriaRepository := repository.NewCategoriaRepository(db)
	if videos, err = categoriaRepository.VideosBelongsToCategoria(uint(IDUint), videosWithPagination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videosWithPagination)
}
