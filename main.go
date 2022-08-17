package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var patients = createDummyPatients()

func getPatients(context *gin.Context) {
	gender := createPatient()
	fmt.Println(gender)
	context.IndentedJSON(http.StatusOK, patients)
}

func addPatient(context *gin.Context) {
	var p patient

	if err := context.BindJSON(&p); err != nil {
		return
	}

	patients = append(patients, p)
	context.IndentedJSON(http.StatusCreated, patients)

}

func getPatient(context *gin.Context) {
	id := context.Param("id")
	patient, error := getPatientById(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Patient not found"})
	}

	context.IndentedJSON(http.StatusOK, patient)
}

func updatePatient(context *gin.Context) {
	id := context.Param("id")
	patient, error := getPatientById(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Patient not found"})
	}

	patient.Active = !patient.Active
	context.IndentedJSON(http.StatusOK, patient)
}

func getPatientById(id string) (*patient, error) {
	for i, p := range patients {
		if p.ID == id {
			return &patients[i], nil
		}
	}
	return nil, errors.New("patient not found")
}

func main() {
	router := gin.Default()
	router.GET("/patients", getPatients)
	router.GET("/patients/:id", getPatient)
	router.PATCH("/patients/:id", updatePatient)
	router.POST("/patients", addPatient)
	router.Run("localhost:9090")

}
