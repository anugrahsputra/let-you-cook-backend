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
	UpdateCategory(id string, userId string, update dto.ReqCreateCategory) (model.Category, error)
	DeleteCategory(id string, userId string) (model.Category, error)
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
		return errors.New("kategori dengan nama ini sudah ada")
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

func (r *categoryRepo) UpdateCategory(id string, userId string, update dto.ReqCreateCategory) (model.Category, error) {
	collection := r.db.Collection("categories")

	var existingCategory model.Category
	err := collection.FindOne(context.Background(), bson.M{
		"id":      id,
		"user_id": userId,
	}).Decode(&existingCategory)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Category{}, errors.New("category not found")
		}
		return model.Category{}, err
	}

	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"id": id, "user_id": userId},
		bson.M{"$set": update},
	)

	if err != nil {
		return model.Category{}, err
	}

	var updatedCategory model.Category
	if err = collection.FindOne(context.Background(), bson.M{
		"id":      id,
		"user_id": userId,
	}).Decode(&updatedCategory); err != nil {
		return model.Category{}, err
	}

	return updatedCategory, nil

}

func (r *categoryRepo) DeleteCategory(id string, userid string) (model.Category, error) {
	collection := r.db.Collection("categories")

	var categoryToDelete model.Category
	err := collection.FindOne(context.Background(), bson.M{
		"id":      id,
		"user_id": userid,
	}).Decode(&categoryToDelete)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Category{}, errors.New("category not found")
		}
		return model.Category{}, err
	}
	_, err = collection.DeleteOne(context.Background(), bson.M{
		"id":      id,
		"user_id": userid,
	})
	if err != nil {
		return model.Category{}, err
	}

	return categoryToDelete, nil
}
