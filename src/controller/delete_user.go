package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init Delete User controller",
		zap.String("journey", "Delete User"),
	)

	userId := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		rest_err := rest_err.NewBadRequestError("Invalid user ID")
		c.JSON(rest_err.Code, rest_err)
	}

	err := uc.service.DeleteUserService(userId)
	if err != nil {
		logger.Error("Error when trying to Delete User",
			err,
			zap.String("Journey", "Delete User"))

		c.JSON(err.Code, err)
	}

	logger.Info("Delete User controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "Delete User"),
	)

	c.Status(http.StatusOK)
}
