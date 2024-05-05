package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/validation"
	"github.com/matheuszzy/crud-basico-go/src/controller/adapters/in"
	"github.com/matheuszzy/crud-basico-go/src/model"
	"github.com/matheuszzy/crud-basico-go/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init login user controller",
		zap.String("journey", "login user"),
	)

	var userRequest in.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("Journey", "login user"))
		resterr := validation.ValidateUserError(err)

		c.JSON(resterr.Code, resterr)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		logger.Error("Error trying to call login user service",
			err,
			zap.String("journey", "login user"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Login controller executed succesfully",
		zap.String("journey", "login user"),
	)
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
