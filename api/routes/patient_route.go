package routes

import (
	"give-me-one-more-shot/server/api/controllers"

	"github.com/gin-gonic/gin"
)

func PatientRoute(router *gin.Engine) {
	router.GET("/patients", controllers.GetPatients())
	router.GET("/patients/:id", controllers.GetPatient())
	router.PATCH("/patients/:id", controllers.UpdatePatient())
	router.POST("/patients", controllers.AddPatient())
}
