package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rwiteshbera/authentication-go-gin/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("users/", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
}
