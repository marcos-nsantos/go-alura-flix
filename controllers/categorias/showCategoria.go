package categoriaControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"net/http"
	"strconv"
)

// ShowCategoria godoc
// @Summary Lista todas as categorias
// @Description Lista todas as categorias
// @Tags categorias
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Categoria "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /categorias [get]
func ShowCategoria(c *gin.Context) {
	IDCategoria := c.Param("id")
	IDUint, err := strconv.ParseUint(IDCategoria, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Inválido"})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var categoria models.Categoria
	categoria.ID = uint(IDUint)

	categoriaRepository := repository.NewCategoriaRepository(db)
	if err := categoriaRepository.FindByID(&categoria); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada"})
		return
	}

	c.JSON(http.StatusOK, categoria)
}
