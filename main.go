package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/controller"
	"github.com/matheuszzy/crud-basico-go/src/controller/routes"
	"github.com/matheuszzy/crud-basico-go/src/model/service"
)

func main() {
	logger.Info("Starting application...")
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
