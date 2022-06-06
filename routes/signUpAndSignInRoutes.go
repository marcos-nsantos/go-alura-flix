package routes

import (
	"github.com/gin-gonic/gin"
	usersController "github.com/marcos-nsantos/alura-flix/controllers/users"
)

func addSignInAndSignOutRoutes(rg *gin.RouterGroup) {
	rg.POST("/signup", usersController.CreateUser)
	rg.POST("/signin", usersController.LoginUser)
}
