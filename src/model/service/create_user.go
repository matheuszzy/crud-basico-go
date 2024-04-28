package service

import (
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"github.com/matheuszzy/crud-basico-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("Journey", "CreateUser"))

	userDomain.EncryptPassword()
	return nil
}
