package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheuszzy/crud-basico-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/userById/:id", userController.FindUserByID)
	r.GET("/userByEmail/:email", userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)
}
