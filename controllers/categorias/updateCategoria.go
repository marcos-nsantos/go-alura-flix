package categoriaControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/dto"
	"github.com/marcos-nsantos/alura-flix/models"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
	"strconv"
)

func UpdateCategoria(c *gin.Context) {
	IDCategoria := c.Param("id")
	IDUint, err := strconv.ParseUint(IDCategoria, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var categoriaToUpdate dto.CategoriaDTO
	if err := c.BindJSON(&categoriaToUpdate); err != nil {
		errValidationMessagesResponse := utils.GetErrValidationMessageResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{"erro": errValidationMessagesResponse})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	var categoria models.Categoria
	categoria.ID = uint(IDUint)

	categoriaRepository := repository.NewCategoriaRepository(db)
	if err := categoriaRepository.FindByID(&categoria); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Categoria não encontrada"})
		return
	}

	err = categoriaRepository.UpdateCategoria(&categoria, categoriaToUpdate.ToCategoria())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categoria)
}
