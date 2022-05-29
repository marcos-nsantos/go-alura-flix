package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequests() {
	r := gin.Default()

	video := r.Group("/videos")
	addVideoRoutes(video)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
