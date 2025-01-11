package repository

import (
	"context"
	"errors"
	"let-you-cook/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITaskRepo interface {
	CreateTask(task model.Task) error
	GetTasks(userId string) ([]model.Task, error)
	GetTaskGroupedByCategory(userId string) ([]model.TaskByCategoryGroup, error)
	UpdateTask(id string, userId string, update map[string]interface{}) (model.Task, error)
	DeleteTask(id string, userId string) (model.Task, error)
}

type taskRepo struct {
	db        *mongo.Database
	indexRepo *IndexRepo
}

func NewTaskRepo(db *mongo.Database, indexRepo *IndexRepo) *taskRepo {
	return &taskRepo{
		db:        db,
		indexRepo: indexRepo,
	}
}

func (r *taskRepo) CreateTask(task model.Task) error {
	collection := r.db.Collection("tasks")

	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		return err
	}
	return nil
}

func (r *taskRepo) GetTasks(userId string) ([]model.Task, error) {
	collection := r.db.Collection("tasks")

	var tasks []model.Task
	cursor, err := collection.Find(context.Background(), bson.M{"user_id": userId})
	if err != nil {
		return nil, errors.New("failed to fetch tasks")
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task model.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepo) GetTaskGroupedByCategory(userId string) ([]model.TaskByCategoryGroup, error) {
	collection := r.db.Collection("tasks")

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.M{"user_id": userId}},
		},
		{
			{Key: "$lookup", Value: bson.M{
				"from":         "categories",
				"localField":   "category_id",
				"foreignField": "id",
				"as":           "category",
			}},
		},
		{
			{Key: "$unwind", Value: bson.M{
				"path":                       "$category",
				"preserveNullAndEmptyArrays": false,
			}},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$category.id",
				"category": bson.M{
					"$first": "$category",
				},
				"tasks": bson.M{
					"$push": bson.M{
						"id":           "$id",
						"user_id":      "$user_id",
						"title":        "$title",
						"description":  "$description",
						"status":       "$status",
						"priority":     "$priority",
						"created_at":   "$created_at",
						"updated_at":   "$updated_at",
						"completed_at": "$completed_at",
						"tags":         "$tags",
					},
				},
			}},
		},
		{
			{Key: "$project", Value: bson.M{
				"category_id": "$_id",
				"category":    1,
				"tasks":       1,
				"_id":         0,
			}},
		},
	}

	var result []model.TaskByCategoryGroup
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var group model.TaskByCategoryGroup
		if err := cursor.Decode(&group); err != nil {
			return nil, err
		}
		result = append(result, group)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *taskRepo) UpdateTask(id string, userId string, update map[string]interface{}) (model.Task, error) {
	collection := r.db.Collection("tasks")

	var existingTask model.Task
	err := collection.FindOne(context.Background(), bson.M{
		"id":      id,
		"user_id": userId,
	}).Decode(&existingTask)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Task{}, errors.New("task not found or unauthorized")
		}
		return model.Task{}, err
	}

	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"id": id, "user_id": userId},
		bson.M{"$set": update},
	)
	if err != nil {
		return model.Task{}, err
	}

	var updatedTask model.Task
	if err = collection.FindOne(context.Background(), bson.M{
		"id":      id,
		"user_id": userId,
	}).Decode(&updatedTask); err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil
}

func (r *taskRepo) DeleteTask(id string, userId string) (model.Task, error) {
	collection := r.db.Collection("tasks")

	var taskToDelete model.Task
	err := collection.FindOne(context.Background(), bson.M{
		"id":      id,
		"user_id": userId,
	}).Decode(&taskToDelete)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Task{}, errors.New("task not found")
		}
		return model.Task{}, err
	}
	_, err = collection.DeleteOne(context.Background(), bson.M{
		"id":      id,
		"user_id": userId,
	})
	if err != nil {
		return model.Task{}, err
	}

	return taskToDelete, nil
}
