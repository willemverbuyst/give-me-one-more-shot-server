package responses

import "give-me-one-more-shot/server/api/models"

type UserResponse struct {
	Status  int         `json:"status" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    models.User `json:"data"`
}

type UsersResponse struct {
	Status  int           `json:"status" example:"200"`
	Message string        `json:"message" example:"success"`
	Data    []models.User `json:"data"`
}
