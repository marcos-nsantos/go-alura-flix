package routes

import (
	"github.com/gin-gonic/gin"
	usersController "github.com/marcos-nsantos/alura-flix/controllers/users"
)

func addUserRoutes(rg *gin.RouterGroup) {
	rg.GET("/", usersController.ShowAllUsers)
	rg.GET("/:id", usersController.ShowUser)
	rg.PUT("/:id", usersController.UpdateUser)
	rg.PATCH("/:id", usersController.UpdateUserPassword)
	rg.DELETE("/:id", usersController.DeleteUser)
}
