package converter

import (
	"github.com/matheuszzy/crud-basico-go/src/model"
	"github.com/matheuszzy/crud-basico-go/src/model/repository/entity"
)

func ConverEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetId(entity.ID.Hex())
	return domain
}
