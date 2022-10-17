package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rwiteshbera/authentication-go-gin/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8000"
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Access granted"})
	})

	router.Run(":" + PORT)
}
