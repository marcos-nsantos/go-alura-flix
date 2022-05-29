package routes

import (
	"github.com/gin-gonic/gin"
	videos_controller "github.com/marcos-nsantos/alura-flix/controllers/videos"
)

func addVideoRoutes(rg *gin.RouterGroup) {
	rg.POST("/", videos_controller.CreateVideo)
	rg.GET("/", videos_controller.ShowAllVideos)
}
