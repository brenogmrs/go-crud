package main

import (
	"log"

	"github.com/brenogmrs/go-crud/src/configuration/logger"
	"github.com/brenogmrs/go-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting Application")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
