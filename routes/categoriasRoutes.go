package routes

import (
	"github.com/gin-gonic/gin"
	categoriaControllers "github.com/marcos-nsantos/alura-flix/controllers/categorias"
)

func addCategoriaRoutes(rg *gin.RouterGroup) {
	rg.POST("/", categoriaControllers.CreateCategoria)
}
