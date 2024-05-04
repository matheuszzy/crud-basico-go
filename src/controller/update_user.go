package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"github.com/matheuszzy/crud-basico-go/src/config/validation"
	"github.com/matheuszzy/crud-basico-go/src/controller/adapters/in"
	"github.com/matheuszzy/crud-basico-go/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init Update user controller",
		zap.String("journey", "Update user"),
	)

	var userRequest in.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("Journey", "Update user"))
		resterr := validation.ValidateUserError(err)

		c.JSON(resterr.Code, resterr)
		return
	}

	userId := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		rest_err := rest_err.NewBadRequestError("Invalid user ID")
		c.JSON(rest_err.Code, rest_err)
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUserService(userId, domain)
	if err != nil {
		logger.Error("Error when trying to update user",
			err,
			zap.String("Journey", "Update user"))

		c.JSON(err.Code, err)
	}

	logger.Info("Update user controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "Update user"),
	)

	c.Status(http.StatusOK)
}
