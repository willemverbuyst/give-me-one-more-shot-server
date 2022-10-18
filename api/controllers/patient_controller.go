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

// GetPatients godoc
// @Summary      Get patients array
// @Description  Responds with the list of all patients as JSON.
// @Tags         patients
// @Produce      json
// @Success      200 {object} responses.PatientsResponse
// @Router       /patients [get]
func GetPatients() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, responses.PatientsResponse{Status: http.StatusOK, Message: "success", Data: patients})
	}
}

// AddPatient godoc
// @Summary      Add patient
// @Description  Responds with the list of all patients as JSON.
// @Tags         patients
// @Accept       json
// @Produce      json
// @Param        patient body models.Patient true "Add patient"
// @Success      200 {object} responses.PatientsResponse
// @Router       /patients [post]
func AddPatient() gin.HandlerFunc {
	return func(context *gin.Context) {
		var p models.Patient

		if err := context.BindJSON(&p); err != nil {
			return
		}

		patients = append(patients, p)
		context.IndentedJSON(http.StatusCreated, responses.PatientsResponse{Status: http.StatusOK, Message: "success", Data: patients})
	}
}

// GetPatient godoc
// @Summary      Get patient by ID
// @Description  Responds with a patient as JSON
// @Tags         patients
// @Accept       json
// @Produce      json
// @Param        id path string true "Patient ID"
// @Success      200 {object} responses.PatientResponse
// @Failure      404 {object} responses.HTTPError
// @Router       /patients/{id} [get]
func GetPatient() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		index := helpers.GetPatientById(id, patients)

		if index == -1 {
			context.IndentedJSON(http.StatusNotFound, responses.HTTPError{Status: http.StatusNotFound, Message: "fail"})
		}

		patient := patients[index]
		context.IndentedJSON(http.StatusOK, responses.PatientResponse{Status: http.StatusOK, Message: "success", Data: patient})
	}
}

// UpdatePatient godoc
// @Summary      Toggles active status of patient
// @Description  Responds with a patient as JSON
// @Tags         patients
// @Accept       json
// @Produce      json
// @Param        id path string true "Patient ID"
// @Success      200 {object} responses.PatientResponse
// @Failure      404 {object} responses.HTTPError
// @Router       /patients/{id} [patch]
func UpdatePatient() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		index := helpers.GetPatientById(id, patients)

		if index == -1 {
			context.IndentedJSON(http.StatusNotFound, responses.HTTPError{Status: http.StatusNotFound, Message: "fail"})
		}

		patient := patients[index]
		patient.Active = !patient.Active
		context.IndentedJSON(http.StatusOK, responses.PatientResponse{Status: http.StatusOK, Message: "success", Data: patient})
	}
}
