package repository

import (
	"context"
	"os"

	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"github.com/matheuszzy/crud-basico-go/src/model"
	"github.com/matheuszzy/crud-basico-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init Update User repository", zap.String("Journey", "Update User"))
	collection_name := os.Getenv(COLLECTION_USER)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	filter := bson.M{"_id": userId}

	_, err := collection.UpdateOne(context.Background(), filter, value)
	if err != nil {
		logger.Error("Error trying to update user",
			err,
			zap.String("Journey", "Update User"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("update user service execute successfully",
		zap.String("Journey", "Update User"))

	return nil
}
