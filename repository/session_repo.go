package repository

import (
	"context"
	"errors"
	"let-you-cook/domain/model"
	"let-you-cook/utils/helper"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "pomodoro_sessions"
	fieldId        = "id"
	fieldUserId    = "user_id"
)

type ISessionRepo interface {
	CreateSession(session model.PomodoroSession) error
	StartSession(id string, userId string) error
	EndSession(id string, userId string) error
	GetAllSessions(useId string) ([]model.PomodoroSession, error)
}

type sessionRepo struct {
	db        *mongo.Database
	indexRepo *IndexRepo
}

func NewSessionRepo(db *mongo.Database, indexRepo *IndexRepo) *sessionRepo {
	return &sessionRepo{
		db:        db,
		indexRepo: indexRepo,
	}
}

func (r *sessionRepo) CreateSession(session model.PomodoroSession) error {
	if err := helper.ValidateSession(session); err != nil {
		return err
	}

	collection := r.db.Collection(collectionName)
	var existingSession model.PomodoroSession
	err := collection.FindOne(
		context.Background(),
		bson.M{"name": bson.M{"$regex": "^" + regexp.QuoteMeta(session.Name) + "$", "$options": "i"}},
	).Decode(&existingSession)

	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	if err == nil {
		return errors.New("session with this name already exists")
	}

	_, err = collection.InsertOne(context.Background(), session)
	if err != nil {
		return err
	}

	return nil
}

func (r *sessionRepo) StartSession(id string, userId string) error {
	collection := r.db.Collection(collectionName)

	filter := bson.M{fieldId: id, fieldUserId: userId}
	update := bson.M{"$set": bson.M{"status": "ACTIVE", "start_time": int(time.Now().Unix()), "updated_at": int(time.Now().Unix())}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("session not found")
		}
		return err
	}

	return nil
}

func (r *sessionRepo) EndSession(id string, userId string) error {
	collection := r.db.Collection(collectionName)

	filter := bson.M{fieldId: id, fieldUserId: userId}

	update := bson.M{
		"$set": bson.M{
			"status":     "COMPLETED",
			"end_time":   int(time.Now().Unix()),
			"updated_at": int(time.Now().Unix()),
		},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("session not found")
		}
		return err
	}

	return nil
}

func (r *sessionRepo) GetAllSessions(userId string) ([]model.PomodoroSession, error) {
	collection := r.db.Collection(collectionName)

	var sessions []model.PomodoroSession
	cursor, err := collection.Find(context.Background(), bson.M{fieldUserId: userId})
	if err != nil {
		return nil, errors.New("failed to fetch tasks")
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var session model.PomodoroSession
		if err := cursor.Decode(&session); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return sessions, nil
}
