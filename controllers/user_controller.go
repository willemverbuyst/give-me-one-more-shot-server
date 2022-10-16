package controllers

import (
	"give-me-one-more-shot/server/helpers"
	"give-me-one-more-shot/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, users)
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
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		}

		context.IndentedJSON(http.StatusOK, user)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		user, error := helpers.GetUserById(id, users)

		if error != nil {
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		}

		user.Active = !user.Active
		context.IndentedJSON(http.StatusOK, user)
	}
}
