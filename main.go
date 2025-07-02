package main

import (
	"log"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	api := router.Group("/api")
	routes.InitUserRoutes(api.Group("/user"))

	if err := router.Run(":8080"); err != nil{
		log.Fatal()
	}
}
