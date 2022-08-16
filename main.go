package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type patient struct {
	Birthdate  string `json:"birthdate"`
	BSN        string `json:"BSN"`
	Email      string `json:"email"`
	FamilyName string `json:"familyName"`
	Gender     string `json:"gender"`
	GivenName  string `json:"givenName"`
}

var patients = []patient{{
	Birthdate:  "02/09/2011",
	BSN:        "123123123",
	Email:      "a@a.io",
	FamilyName: "Hay",
	Gender:     "Female",
	GivenName:  "Wong",
},
	{
		Birthdate:  "03/11/1998",
		BSN:        "12456789",
		Email:      "b@b.io",
		FamilyName: "Foo",
		Gender:     "Male",
		GivenName:  "Bar",
	}}

func getPatients(context *gin.Context) {
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

func main() {
	router := gin.Default()
	router.GET("/patients", getPatients)
	router.POST("/patients", addPatient)
	router.Run("localhost:9090")

}
