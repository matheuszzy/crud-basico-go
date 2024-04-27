package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/validation"
	"github.com/matheuszzy/crud-basico-go/src/controller/adapters/in"
	"github.com/matheuszzy/crud-basico-go/src/model"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest in.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("Journey", "CreateUser"))
		resterr := validation.ValidateUserError(err)

		c.JSON(resterr.Code, resterr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.String(http.StatusOK, "")
}
