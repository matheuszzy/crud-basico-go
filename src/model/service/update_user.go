package service

import (
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"github.com/matheuszzy/crud-basico-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init update user model", zap.String("Journey", "update user"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		return err
	}

	logger.Info("update user service execute successfully",
		zap.String("Journey", "update user"))

	return nil
}
