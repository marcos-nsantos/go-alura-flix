package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/alura-flix/security"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := security.ValidateToken(c); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}
