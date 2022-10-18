package responses

import "give-me-one-more-shot/server/api/models"

type PatientResponse struct {
	Status  int            `json:"status" example:"200"`
	Message string         `json:"message" example:"success"`
	Data    models.Patient `json:"data"`
}

type PatientsResponse struct {
	Status  int              `json:"status" example:"200"`
	Message string           `json:"message" example:"success"`
	Data    []models.Patient `json:"data"`
}
