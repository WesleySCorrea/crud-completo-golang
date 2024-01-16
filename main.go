package main

import (
	"log"

	"github.com/WesleySCorrea/crud-completo-golang/src/configuration/logger"
	"github.com/WesleySCorrea/crud-completo-golang/src/controller"
	"github.com/WesleySCorrea/crud-completo-golang/src/controller/routes"
	"github.com/WesleySCorrea/crud-completo-golang/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Start application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
