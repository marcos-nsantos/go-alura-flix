package routes

import (
	"github.com/gin-gonic/gin"
	usersController "github.com/marcos-nsantos/alura-flix/controllers/users"
)

func addUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/", usersController.CreateUser)
	rg.GET("/", usersController.ShowAllUsers)
	rg.GET("/:id", usersController.ShowUser)
}
