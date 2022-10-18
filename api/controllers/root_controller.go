package controllers

import (
	"give-me-one-more-shot/server/api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WelcomeWorld godoc
// @Summary      Test root
// @Description  Responds with "Hello world" message.
// @Tags         root
// @Produce      json
// @Success      200 {object} responses.RootResponse
// @Router       / [get]
func WelcomeWorld() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, responses.RootResponse{Status: http.StatusOK, Message: "Hello World"})
	}
}
