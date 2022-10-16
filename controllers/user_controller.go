package controllers

import (
	"give-me-one-more-shot/server/helpers"
	"give-me-one-more-shot/server/models"
	"give-me-one-more-shot/server/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "All users", Data: nil})
	}
}

func AddUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var u models.User

		if err := context.BindJSON(&u); err != nil {
			return
		}

		users = append(users, u)
		context.IndentedJSON(http.StatusCreated, users)
	}
}

func GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		user, error := helpers.GetUserById(id, users)

		if error != nil {
			context.IndentedJSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "User not found", Data: nil})
		}

		context.IndentedJSON(http.StatusOK, user)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		user, error := helpers.GetUserById(id, users)

		if error != nil {
			context.IndentedJSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "User not found", Data: nil})
		}

		user.Active = !user.Active
		context.IndentedJSON(http.StatusOK, user)
	}
}
