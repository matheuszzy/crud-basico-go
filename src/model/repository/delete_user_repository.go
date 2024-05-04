package repository

import (
	"context"
	"os"

	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"github.com/matheuszzy/crud-basico-go/src/config/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init Delete user repository",
		zap.String("Journey", "Delete user"))

	collection_name := os.Getenv(COLLECTION_USER)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to Delete user",
			err,
			zap.String("Journey", "Delete user"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("Delete user repository execute successfully",
		zap.String("UserId", userId),
		zap.String("Journey", "Delete user"))

	return nil
}
