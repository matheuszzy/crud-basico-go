package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheuszzy/crud-basico-go/src/model/service"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	CreateUser(c *gin.Context)

	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
