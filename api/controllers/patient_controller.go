package controllers

import (
	"give-me-one-more-shot/server/api/helpers"
	"give-me-one-more-shot/server/api/models"
	"give-me-one-more-shot/server/api/responses"

	"net/http"

	"github.com/gin-gonic/gin"
)

var users = helpers.GetUsers()
var patients = helpers.CreateDummyPatients(users)

func GetPatients() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "All patients", Data: patients})
	}
}

func AddPatient() gin.HandlerFunc {
	return func(context *gin.Context) {
		var p models.Patient

		if err := context.BindJSON(&p); err != nil {
			return
		}

		patients = append(patients, p)
		context.IndentedJSON(http.StatusCreated, responses.UserResponse{Status: http.StatusOK, Message: "Patient created", Data: patients})
	}
}

func GetPatient() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		patient, error := helpers.GetPatientById(id, patients)

		if error != nil {
			context.IndentedJSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Patient not found", Data: nil})
		}

		context.IndentedJSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "One patient", Data: patient})
	}
}

func UpdatePatient() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		patient, error := helpers.GetPatientById(id, patients)

		if error != nil {
			context.IndentedJSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Patient not found", Data: nil})
		}

		patient.Active = !patient.Active
		context.IndentedJSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "Patient updated", Data: patient})
	}
}
