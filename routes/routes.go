package routes

import (
	"github.com/gin-gonic/gin"
)

func HandleRequests() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	addUserRoutes(users)

	video := r.Group("/videos")
	addVideoRoutes(video)

	categoria := r.Group("/categorias")
	addCategoriaRoutes(categoria)

	return r
}
