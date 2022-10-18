package controllers

import (
	"give-me-one-more-shot/server/api/helpers"
	"give-me-one-more-shot/server/api/models"
	"give-me-one-more-shot/server/api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary      Get users array
// @Description  Responds with the list of all users as JSON.
// @Tags         users
// @Produce      json
// @Success      200 {object} responses.UsersResponse
// @Router       /users [get]
func GetUsers() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, responses.UsersResponse{Status: http.StatusOK, Message: "All users", Data: users})
	}
}

// AddUser godoc
// @Summary      Add User
// @Description  Responds with the list of all users as JSON.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.User true "Add user"
// @Success      200 {object} responses.UsersResponse
// @Router       /users [post]
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

// GetUser godoc
// @Summary      Get user by ID
// @Description  Responds with a user as JSON
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} responses.UserResponse
// @Failure      404 {object} responses.HTTPError
// @Router       /users/{id} [get]
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

// UpdateUser godoc
// @Summary      Toggles active status of user
// @Description  Responds with a user as JSON
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} responses.UserResponse
// @Failure      404 {object} responses.HTTPError
// @Router       /users/{id} [patch]
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
