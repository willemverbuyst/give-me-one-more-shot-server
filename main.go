package main

import (
	"give-me-one-more-shot/server/config"
	"give-me-one-more-shot/server/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.Database()
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
	routes.UserRoute(router)
	routes.RootRoute(router)

	sErr := router.Run(":9090")
	if sErr != nil {
		log.Fatal("Error when running the server:", sErr)
	}
}
