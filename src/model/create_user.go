package model

import (
	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("Journey", "CreateUser"))

	ud.EncryptPassword()

	return nil
}
