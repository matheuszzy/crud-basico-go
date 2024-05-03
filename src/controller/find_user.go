package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"github.com/matheuszzy/crud-basico-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("FindUserByID controller",
		zap.String("Journey", "FindUser"))

	userId := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Info("Error trying to validate user id",
			zap.String("Journey", "FindUser"))
		errorMessage := rest_err.NewBadRequestError("User ID is not valid id")

		c.JSON(errorMessage.Code, errorMessage)
	}

	userDomain, err := uc.service.FindUserByIDService(userId)
	if err != nil {
		logger.Info("Error when finding user by id",
			zap.String("Journey", "FindUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User found successfully",
		zap.String("Journey", "FindUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("FindUserByEmail controller",
		zap.String("Journey", "FindUser"))

	userEmail := c.Param("email")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Info("Error trying to validate user email",
			zap.String("Journey", "FindUser"))
		errorMessage := rest_err.NewBadRequestError("User email is not a valid email")

		c.JSON(errorMessage.Code, errorMessage)
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)
	if err != nil {
		logger.Info("Error when finding user by email",
			zap.String("Journey", "FindUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User found successfully",
		zap.String("Journey", "FindUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
