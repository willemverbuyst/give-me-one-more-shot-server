package controllers

import (
	"give-me-one-more-shot/server/api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeWorld() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "Hello World", Data: nil})
	}
}
