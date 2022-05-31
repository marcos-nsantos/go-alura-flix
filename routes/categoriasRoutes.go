package routes

import (
	"github.com/gin-gonic/gin"
	categoriaControllers "github.com/marcos-nsantos/alura-flix/controllers/categorias"
)

func addCategoriaRoutes(rg *gin.RouterGroup) {
	rg.POST("/", categoriaControllers.CreateCategoria)
	rg.GET("/", categoriaControllers.ShowAllCategorias)
	rg.GET("/:id", categoriaControllers.ShowCategoria)
	rg.PUT("/:id", categoriaControllers.UpdateCategoria)
	rg.DELETE("/:id", categoriaControllers.DeleteCategoria)
	rg.GET("/:id/videos", categoriaControllers.ShowVideosByCategoria)
}
