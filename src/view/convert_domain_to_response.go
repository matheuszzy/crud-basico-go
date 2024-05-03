package view

import (
	"github.com/matheuszzy/crud-basico-go/src/controller/adapters/out"
	"github.com/matheuszzy/crud-basico-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) out.UserResponse {
	return out.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
