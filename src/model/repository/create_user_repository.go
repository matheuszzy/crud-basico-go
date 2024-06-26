package repository

import (
	"context"
	"os"

	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"github.com/matheuszzy/crud-basico-go/src/model"
	"github.com/matheuszzy/crud-basico-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser repository", zap.String("Journey", "CreateUser"))

	collection_name := os.Getenv(COLLECTION_USER)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error(err.Error(),
			err,
			zap.String("Journey", "Update User"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("update user service execute successfully",
		zap.Any("userID", value.ID.Hex()),
		zap.String("Journey", "update user"))

	return converter.ConverEntityToDomain(*value), nil
}
