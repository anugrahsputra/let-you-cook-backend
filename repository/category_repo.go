package repository

import (
	"context"
	"errors"
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICategoryRepo interface {
	CreateCategory(category model.Category) error
	GetCategories(userId string, reqCategory dto.ReqCategory) ([]model.Category, error)
	GetCategoryById(id string, userId string) (model.Category, error)
	UpdateCategory(id string, userId string, category model.Category) error
	DeleteCategory(id string, userId string) error
}

type categoryRepo struct {
	db        *mongo.Database
	indexRepo *IndexRepo
}

func NewCategoryRepo(db *mongo.Database, indexRepo *IndexRepo) *categoryRepo {
	return &categoryRepo{
		db:        db,
		indexRepo: indexRepo,
	}
}

func (r *categoryRepo) CreateCategory(category model.Category) error {
	collection := r.db.Collection("categories")

	var existingCategory model.Category
	err := collection.FindOne(context.Background(), bson.M{"name": bson.M{"$regex": "^" + regexp.QuoteMeta(category.Name) + "$", "$options": "i"}}).Decode(&existingCategory)
	if err == nil {
		return errors.New("category with this name already exists")
	}
	if err != mongo.ErrNoDocuments {
		return err
	}
	_, err = collection.InsertOne(context.Background(), category)
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryRepo) GetCategories(userId string, reqCategory dto.ReqCategory) ([]model.Category, error) {
	var categories []model.Category

	filter := bson.M{"user_id": userId}
	if reqCategory.Id != "" {
		filter["id"] = reqCategory.Id
	}

	if reqCategory.Name != "" {
		filter["name"] = bson.M{
			"$regex":   `(?i)` + reqCategory.Name,
			"$options": "i",
		}
	}

	collection := r.db.Collection("categories")
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.New("failed to fetch categories")
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var category model.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepo) GetCategoryById(id string, userId string) (model.Category, error) {
	collection := r.db.Collection("categories")

	var category model.Category
	err := collection.FindOne(
		context.Background(),
		bson.M{"id": id, "user_id": userId},
	).Decode(&category)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return category, nil
		}
		return model.Category{}, err
	}
	return category, nil
}

func (r *categoryRepo) UpdateCategory(id string, userId string, update model.Category) error {
	collection := r.db.Collection("categories")

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"id": id, "user_id": userId},
		bson.M{"$set": update},
	)

	if err != nil {
		return err
	}

	return nil

}

func (r *categoryRepo) DeleteCategory(id string, userid string) error {
	collection := r.db.Collection("categories")

	_, err := collection.DeleteOne(context.Background(), bson.M{
		"id":      id,
		"user_id": userid,
	})
	if err != nil {
		return err
	}

	return nil
}
