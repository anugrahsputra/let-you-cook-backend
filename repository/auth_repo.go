package repository

import (
	"context"
	"errors"
	"let-you-cook/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthRepo interface {
	RegisterRepo(uesr model.User) error
	CheckUserExistingForLogin(username string) (model.User, error)
	GetUserExisting(username string) (model.User, error)
}

type authRepo struct {
	db        *mongo.Database
	indexRepo *IndexRepo
}

func NewAuthRepository(db *mongo.Database, indexRepo *IndexRepo) *authRepo {
	err := indexRepo.CreateIndex("users", "username", true)
	if err != nil {
		logger.Error("error creating index: %v\n", err)
	}
	return &authRepo{
		db:        db,
		indexRepo: indexRepo,
	}
}

func (r *authRepo) RegisterRepo(user model.User) error {
	collection := r.db.Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("user already exist")
		}
		return err
	}
	return nil
}

func (r *authRepo) CheckUserExistingForLogin(username string) (model.User, error) {
	var user model.User
	filter := bson.M{"username": username}

	err := r.db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNilDocument {
			return user, errors.New("user not found")
		}
		logger.Errorf("error finding user: %v\n", err)
		return user, err
	}
	return user, nil
}

func (r *authRepo) GetUserExisting(username string) (model.User, error) {
	collection := r.db.Collection("users")
	var user model.User

	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, nil
		}
		return model.User{}, err
	}
	return model.User{}, nil
}
