package routes

import (
	"give-me-one-more-shot/server/controllers"

	"github.com/gin-gonic/gin"
)

func RootRoute(router *gin.Engine) {
	router.GET("/", controllers.WelcomeWorld())
}
