package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/validation"
	"github.com/matheuszzy/crud-basico-go/src/controller/adapters/in"
	"github.com/matheuszzy/crud-basico-go/src/model"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init Update user controller",
		zap.String("journey", "Update user"),
	)

	var userRequest in.UserUpdateRequest
	userId := c.Param("id")

	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validate user info", err,
			zap.String("Journey", "Update user"))
		resterr := validation.ValidateUserError(err)

		c.JSON(resterr.Code, resterr)
		return
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
