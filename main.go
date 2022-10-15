package main

import (
	"log"

	"give-me-one-more-shot/server/models"
	"give-me-one-more-shot/server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := models.Database()
	if err != nil {
		log.Fatalln(err)
	}
	_, dbErr := db.DB()
	if dbErr != nil {
		log.Fatal("Error when starting the database:", err)
	}

	router := gin.Default()
	router.Use(cors.Default())
	routes.PatientRoute(router)
	routes.RootRoute(router)

	sErr := router.Run(":9090")
	if sErr != nil {
		log.Fatal("Error when running the server:", err)
	}
}
