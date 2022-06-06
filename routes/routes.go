package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/middlewares"
)

func HandleRequests() *gin.Engine {
	r := gin.Default()

	signInAndSignOutRoutes := r.Group("/")
	addSignInAndSignOutRoutes(signInAndSignOutRoutes)

	users := r.Group("/users")
	users.Use(middlewares.Authentication())
	addUserRoutes(users)

	videos := r.Group("/videos")
	videos.Use(middlewares.Authentication())
	addVideoRoutes(videos)

	categorias := r.Group("/categorias")
	categorias.Use(middlewares.Authentication())
	addCategoriaRoutes(categorias)

	return r
}
