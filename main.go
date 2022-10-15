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
	db.DB()

	router := gin.Default()
	router.Use(cors.Default())
	routes.PatientRoute(router)
	routes.RootRoute(router)

	router.Run(":9090")
}
