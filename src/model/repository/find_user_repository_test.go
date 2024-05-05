package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/matheuszzy/crud-basico-go/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_vallid_email_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Age:      90,
			Name:     "test",
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
	})

	mtestDb.Run("returns when mongo db returns error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns no documents found", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Age:      90,
			Name:     "test",
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
	})
}

func TestUserRepository_FindUserById(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_vallid_id_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Age:      90,
			Name:     "test",
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
	})

	mtestDb.Run("returns when mongo db returns error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID("test")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns no documents found", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Age:      90,
			Name:     "test",
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_vallid_email_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Age:      90,
			Name:     "test",
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
	})

	mtestDb.Run("returns when mongo db returns error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("test", "123")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns no documents found", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Age:      90,
			Name:     "test",
			Email:    "test@test.com",
			Password: "test",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
