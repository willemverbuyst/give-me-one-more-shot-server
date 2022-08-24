package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/willemverbuyst/give-me-one-more-shot/server/message"
	"github.com/willemverbuyst/give-me-one-more-shot/server/model"
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
	db, err := model.Database()
	if err != nil {
		log.Fatalln(err)
	}
	db.DB()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", welcomeWorld)
	router.GET("/messages", message.GetMessages)
	router.POST("messages", message.PostMessage)
	router.GET("/patients", getPatients)
	router.GET("/patients/:id", getPatient)
	router.PATCH("/patients/:id", updatePatient)
	router.POST("/patients", addPatient)
	router.Run(":9090")
}
