package repository

import (
	"context"
	"fmt"

	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"github.com/matheuszzy/crud-basico-go/src/model"
	"github.com/matheuszzy/crud-basico-go/src/model/repository/entity"
	"github.com/matheuszzy/crud-basico-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository",
		zap.String("Journey", "findUserByEmail"))

	collection := ur.databaseConnection.Collection(COLLECTION_USER)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("Journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("Journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByEmail repository successfully",
		zap.String("Journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))

	return converter.ConverEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository",
		zap.String("Journey", "findUserByID"))

	collection := ur.databaseConnection.Collection(COLLECTION_USER)

	userEntity := &entity.UserEntity{}
	userId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userId}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this id: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("Journey", "findUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage,
			err,
			zap.String("Journey", "findUserByID"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("findUserByID repository successfully",
		zap.String("Journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()))

	return converter.ConverEntityToDomain(*userEntity), nil
}
