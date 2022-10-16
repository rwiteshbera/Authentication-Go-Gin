package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rwiteshbera/authentication-go-gin/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	router := gin.New()

	routes.AuthRoutes(router)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Access granted"})
	})

	router.Run(":" + PORT)
}
