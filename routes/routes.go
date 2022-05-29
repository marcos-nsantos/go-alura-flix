package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequests() {
	r := gin.Default()

	video := r.Group("/videos")
	addVideoRoutes(video)

	categoria := r.Group("/categorias")
	addCategoriaRoutes(categoria)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
