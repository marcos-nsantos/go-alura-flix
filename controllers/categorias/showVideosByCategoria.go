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

// ShowVideosByCategoria godoc
// @Summary Lista todas as categorias
// @Description Lista todas as categorias
// @Tags categorias
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param sort query string false "Sort"
// @Success 200 {array} models.Video "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /categorias [get]
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

	videosPagination := utils.GeneratePagination[models.Video](c)

	categoriaRepository := repository.NewCategoriaRepository(db)
	if err = categoriaRepository.VideosBelongsToCategoria(uint(IDUint), &videosPagination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videosPagination)
}
