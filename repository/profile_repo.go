package repository

import (
	"context"
	"errors"
	"let-you-cook/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProfileRepo interface {
	CreateProfile(profile model.Profile) error
	GetProfileByAccountId(userId string) (model.Profile, error)
	UpdateProfile(userId string, payload model.Profile) error
}

type profileRepo struct {
	db        *mongo.Database
	indexRepo *IndexRepo
}

func NewProfileRepo(db *mongo.Database, indexRepo *IndexRepo) *profileRepo {
	return &profileRepo{
		db:        db,
		indexRepo: indexRepo,
	}
}

func (r *profileRepo) CreateProfile(profile model.Profile) error {
	collection := r.db.Collection("profiles")

	_, err := collection.InsertOne(context.Background(), profile)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("profile already exists")
		}
		return err
	}
	return nil
}

func (r *profileRepo) GetProfileByAccountId(userId string) (model.Profile, error) {
	collection := r.db.Collection("profiles")

	var profile model.Profile
	err := collection.FindOne(context.Background(), bson.M{"user_id": userId}).Decode(&profile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Profile{}, errors.New("profile not found")
		}
		return model.Profile{}, err
	}
	return profile, nil
}

func (r *profileRepo) UpdateProfile(userId string, payload model.Profile) error {
	collection := r.db.Collection("profiles")

	filter := bson.M{"user_id": userId}
	update := bson.M{"$set": payload}

	_, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return err
	}

	return nil
}
