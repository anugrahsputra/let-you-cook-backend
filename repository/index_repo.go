package repository

import (
	"context"

	"github.com/op/go-logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger = logging.MustGetLogger("main")

type IndexRepo struct {
	db *mongo.Database
}

func NewIndexRepo(db *mongo.Database) *IndexRepo {
	return &IndexRepo{
		db: db,
	}
}

func (s *IndexRepo) CreateIndex(collectionName string, field string, unique bool) error {
	collection := s.db.Collection(collectionName)
	indexModel := mongo.IndexModel{
		Keys: bson.M{field: 1}, // 1 untuk ascending, -1 untuk descending
	}

	if unique {
		indexModel.Options = options.Index().SetUnique(true) // Menetapkan indeks sebagai unik jika diperlukan
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return err
	}
	logger.Infof("Index for field '%s' in collection '%s' created successfully.", field, collectionName)
	return nil
}
