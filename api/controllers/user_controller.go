package controllers

import (
	"give-me-one-more-shot/server/api/helpers"
	"give-me-one-more-shot/server/api/models"
	"give-me-one-more-shot/server/api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, responses.UsersResponse{Status: http.StatusOK, Message: "All users", Data: users})
	}
}

func AddUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var u models.User

		if err := context.BindJSON(&u); err != nil {
			return
		}

		users = append(users, u)
		context.IndentedJSON(http.StatusCreated, responses.UsersResponse{Status: http.StatusOK, Message: "User created", Data: users})
	}
}

func GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		index := helpers.GetUserById(id, users)

		if index == -1 {
			context.IndentedJSON(http.StatusNotFound, responses.HTTPError{Status: http.StatusNotFound, Message: "User not found"})
		}

		user := users[index]
		context.IndentedJSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "One user", Data: user})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		index := helpers.GetUserById(id, users)

		if index == -1 {
			context.IndentedJSON(http.StatusNotFound, responses.HTTPError{Status: http.StatusNotFound, Message: "User not found"})
		}

		user := users[index]
		user.Active = !user.Active
		context.IndentedJSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "User updated", Data: user})
	}
}
