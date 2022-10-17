package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/authentication-go-gin/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no authorization header provided"})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
		}

		c.Set("email", claims.Email)
		c.Set("firstname", claims.FirstName)
		c.Set("lastname", claims.LastName)
		c.Set("uid", claims.UserId)
		c.Next()
	}
}
