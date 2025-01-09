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

	// Iterate over the results
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

func (r *taskRepo) UpdateTask(id string, userId string, update map[string]interface{}) (model.Task, error) {
	collection := r.db.Collection("tasks")

	_, err := collection.UpdateOne(context.Background(), bson.M{"id": id}, bson.M{"$set": update})
	if err != nil {
		return model.Task{}, err
	}

	var updatedTask model.Task
	if err = collection.FindOne(context.Background(), bson.M{"user_id": id}).Decode(&updatedTask); err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil
}

func (r *taskRepo) DeleteTask(id string, userId string) (model.Task, error) {
	collection := r.db.Collection("tasks")

	_, err := collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return model.Task{}, err
	}

	return model.Task{}, nil
}
