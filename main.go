package main

import (
	"give-me-one-more-shot/server/api/config"
	"give-me-one-more-shot/server/api/routes"
	_ "give-me-one-more-shot/server/docs"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Give Me One More Shot API
// @version         1.0
// @description     Give Me One More Shot service API in Go using Gin framework.
// @host						localhost:9090
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

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	sErr := router.Run(":9090")
	if sErr != nil {
		log.Fatal("Error when running the server:", sErr)
	}
}
