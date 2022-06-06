package categoriaControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/dto"
	"github.com/marcos-nsantos/alura-flix/repository"
	"github.com/marcos-nsantos/alura-flix/utils"
	"net/http"
)

// CreateCategoria godoc
// @Summary Create categoria
// @Description Create categoria
// @Tags categorias
// @Accept  json
// @Produce  json
// @Param categoria body dto.CategoriaDTO true "Categoria"
// @Success 201 {object} models.Categoria
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /categorias [post]
func CreateCategoria(c *gin.Context) {
	var categoriaToCreate dto.CategoriaDTO

	if err := c.BindJSON(&categoriaToCreate); err != nil {
		errValidationMessagesResponse := utils.GetErrValidationMessageResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errValidationMessagesResponse})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	categoria := categoriaToCreate.ToCategoria()
	categoriaRepository := repository.NewCategoriaRepository(db)
	err = categoriaRepository.CreateCategoria(categoria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, categoria)
}
