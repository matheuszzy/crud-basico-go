package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheuszzy/crud-basico-go/src/controller"
	"github.com/matheuszzy/crud-basico-go/src/model"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/userById/:id", model.IsTokenValid, userController.FindUserByID)
	r.GET("/userByEmail/:email", model.IsTokenValid, userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:id", model.IsTokenValid, userController.UpdateUser)
	r.DELETE("/user/:id", model.IsTokenValid, userController.DeleteUser)

	r.POST("/login", userController.LoginUser)
}
