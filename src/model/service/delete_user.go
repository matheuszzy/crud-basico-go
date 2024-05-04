package service

import (
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init update user model", zap.String("Journey", "Delete user"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "Delete user"))
		return err
	}

	logger.Info("delete user service execute successfully",
		zap.String("Journey", "update user"))

	return nil
}
