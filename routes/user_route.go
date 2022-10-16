package routes

import (
	"give-me-one-more-shot/server/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers())
	router.GET("/users/:id", controllers.GetUser())
	router.PATCH("/users/:id", controllers.UpdateUser())
	router.POST("/users", controllers.AddUser())
}
