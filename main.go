package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var patients = createDummyPatients()

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

func welcomeWorld(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
	db, err := initStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}

	router := gin.Default()
	router.Use(cors.Default())

	defer db.Close()

	router.GET("/", welcomeWorld)
	router.GET("/patients", getPatients)
	router.GET("/patients/:id", getPatient)
	router.PATCH("/patients/:id", updatePatient)
	router.POST("/patients", addPatient)
	router.Run(":9090")
}
