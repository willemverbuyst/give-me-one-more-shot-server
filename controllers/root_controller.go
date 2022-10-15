package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeWorld() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
	}
}
