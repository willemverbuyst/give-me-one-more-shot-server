package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var patients = []patient{{
	ID:         "1A",
	Birthdate:  "02/09/2011",
	BSN:        "123123123",
	Email:      "a@a.io",
	FamilyName: "Hay",
	Gender:     "Female",
	GivenName:  "Wong",
},
	{
		ID:         "2B",
		Birthdate:  "03/11/1998",
		BSN:        "12456789",
		Email:      "b@b.io",
		FamilyName: "Foo",
		Gender:     "Male",
		GivenName:  "Bar",
	}}

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
	router.POST("/patients", addPatient)
	router.Run("localhost:9090")

}
