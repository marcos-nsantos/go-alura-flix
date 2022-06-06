package categoriaControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"net/http"
	"strconv"
)

// DeleteCategoria godoc
// @Summary Delete categoria
// @Description Delete categoria
// @Tags categorias
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /categorias/{id} [delete]
func DeleteCategoria(c *gin.Context) {
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
		c.JSON(http.StatusNotFound, gin.H{"erro": "Categoria não encontrada"})
		return
	}

	if err := categoriaRepository.DeleteCategoria(&categoria); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Categoria deletada com sucesso"})
}
