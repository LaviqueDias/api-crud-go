package main

import (
	"log"

	mysqldb "github.com/LaviqueDias/api-crud-go/src/configuration/database/mysql"
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger.Info("Init Connection to mysql", zap.String("journey", "connectionToDB"),)
	database, err := mysqldb.Connection()
	if err != nil {
		logger.Info("Error connecting to MySQL",
		zap.String("journey", "connectionToDB"),)
		panic(err)
	}
	logger.Info("Connection to mysql successfully", zap.String("journey", "connectionToDB"),)
	defer mysqldb.CloseConnection()
	
	userController := initDependencies(database)

	router := gin.Default()
	api := router.Group("/api")
	routes.InitUserRoutes(api.Group("/user"), userController)

	if err := router.Run(":8080"); err != nil{
		log.Fatal()
	}
}
