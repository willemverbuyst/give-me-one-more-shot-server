package controllers

import (
	"net/http"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
)

var users = helpers.GetUsers()
var patients = helpers.CreateDummyPatients(users)

func GetPatients() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, patients)
	}
}

func AddPatient() gin.HandlerFunc {
	return func(context *gin.Context) {
		var p models.Patient

		if err := context.BindJSON(&p); err != nil {
			return
		}

		patients = append(patients, p)
		context.IndentedJSON(http.StatusCreated, patients)
	}
}

func GetPatient() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		patient, error := helpers.GetPatientById(id, patients)

		if error != nil {
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Patient not found"})
		}

		context.IndentedJSON(http.StatusOK, patient)
	}
}

func UpdatePatient() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		patient, error := helpers.GetPatientById(id, patients)

		if error != nil {
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Patient not found"})
		}

		patient.Active = !patient.Active
		context.IndentedJSON(http.StatusOK, patient)
	}
}
