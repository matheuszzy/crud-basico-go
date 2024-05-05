package service

import (
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"github.com/matheuszzy/crud-basico-go/src/model"
	"go.uber.org/zap"
)

func (ur *userDomainService) FindUserByIDService(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Find user by ID service", zap.String("Journey", "FindUser"))
	return ur.userRepository.FindUserByID(id)
}

func (ur *userDomainService) FindUserByEmailService(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Find user by Email service", zap.String("Journey", "FindUser"))
	return ur.userRepository.FindUserByEmail(email)
}

func (ur *userDomainService) findUserByEmailAndPasswordService(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Find user by email and password service", zap.String("Journey", "FindUser"))
	return ur.userRepository.FindUserByEmailAndPassword(email, password)
}
