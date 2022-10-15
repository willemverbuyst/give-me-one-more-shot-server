package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func RootRoute(router *gin.Engine) {
	router.GET("/", controllers.WelcomeWorld())
}
