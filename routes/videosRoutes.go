package routes

import (
	"github.com/gin-gonic/gin"
	videosController "github.com/marcos-nsantos/alura-flix/controllers/videos"
)

func addVideoRoutes(rg *gin.RouterGroup) {
	rg.POST("/", videosController.CreateVideo)
	rg.GET("/", videosController.ShowAllVideos)
	rg.GET("/:id", videosController.ShowVideo)
	rg.PUT("/:id", videosController.UpdateVideo)
	rg.DELETE("/:id", videosController.DeleteVideo)
}
