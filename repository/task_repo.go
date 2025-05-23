package repository

import (
	"context"
	"errors"
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITaskRepo interface {
	CreateTask(task model.Task) error
	GetTasks(userId string) ([]model.Task, error)
	GetTaskGroupedByCategory(userId string) ([]dto.TaskByCategoryGroupResp, error)
	UpdateTask(id string, userId string, update model.Task) error
	DeleteTask(id string, userId string) error
	FindTask(id string, userId string) (model.Task, error)
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

func (r *taskRepo) GetTaskGroupedByCategory(userId string) ([]dto.TaskByCategoryGroupResp, error) {
	collection := r.db.Collection("categories")

	pipeline := mongo.Pipeline{
		{
			{Key: "$lookup", Value: bson.M{
				"from":         "tasks",
				"localField":   "id",
				"foreignField": "category_id",
				"pipeline": []bson.M{{
					"$match": bson.M{"user_id": userId},
				}},
				"as": "tasks",
			}},
		},
		{
			{Key: "$project", Value: bson.M{
				"category_id":   "$id",
				"category_name": "$name",
				"tasks": bson.M{
					"$filter": bson.M{
						"input": "$tasks",
						"as":    "task",
						"cond":  bson.M{"$ne": bson.A{"$$task", nil}},
					},
				},
				"_id": 0,
			}},
		},
	}

	var result []dto.TaskByCategoryGroupResp
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var group dto.TaskByCategoryGroupResp
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

func (r *taskRepo) UpdateTask(id string, userId string, update model.Task) error {
	collection := r.db.Collection("tasks")

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

func (r *taskRepo) DeleteTask(id string, userId string) error {
	collection := r.db.Collection("tasks")

	_, err := collection.DeleteOne(context.Background(), bson.M{
		"id":      id,
		"user_id": userId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepo) FindTask(id string, userId string) (model.Task, error) {
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

	return existingTask, nil
}
