package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheuszzy/crud-basico-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/userById/:id", controller.FindUserByID)
	r.GET("/userByEmail/:email", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)
}
