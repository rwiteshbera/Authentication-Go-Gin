package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rwiteshbera/authentication-go-gin/controllers"
	middlewares "github.com/rwiteshbera/authentication-go-gin/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.Authenticate()) // Check whether the routes are authenticated or not
	incomingRoutes.GET("users/", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
}
