package routes

import (
	"github.com/gin-gonic/gin"
	videoControllers "github.com/marcos-nsantos/alura-flix/controllers/videos"
	"github.com/marcos-nsantos/alura-flix/docs"
	"github.com/marcos-nsantos/alura-flix/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "Alura Flix"
	docs.SwaggerInfo.Description = "API desenvolvida para o challenge de back-end da Alura"

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

	r.GET("/videos/free/", videoControllers.ShowAllFreeVideos)

	return r
}
