package routes

import (
	"github.com/gin-gonic/gin"
	videoControllers "github.com/marcos-nsantos/alura-flix/controllers/videos"
)

func addVideoRoutes(rg *gin.RouterGroup) {
	rg.POST("/", videoControllers.CreateVideo)
	rg.GET("/", videoControllers.ShowAllVideos)
	rg.GET("/:id", videoControllers.ShowVideo)
	rg.PUT("/:id", videoControllers.UpdateVideo)
	rg.DELETE("/:id", videoControllers.DeleteVideo)
}
