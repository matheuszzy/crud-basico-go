package service

import (
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"github.com/matheuszzy/crud-basico-go/src/model"
)

func (ud *userDomainService) UpdateUserService(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
