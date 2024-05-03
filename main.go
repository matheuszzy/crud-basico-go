package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matheuszzy/crud-basico-go/src/config/database"
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/controller"
	"github.com/matheuszzy/crud-basico-go/src/controller/routes"
	"github.com/matheuszzy/crud-basico-go/src/model/repository"
	"github.com/matheuszzy/crud-basico-go/src/model/service"
	"golang.org/x/net/context"
)

func main() {
	logger.Info("Starting application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database, err := database.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to MongoDB, error=%s", err.Error())
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
