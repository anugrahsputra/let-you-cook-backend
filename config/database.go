package config

import (
	"context"
	"let-you-cook/utils/helper"
	"time"

	"github.com/op/go-logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger = logging.MustGetLogger("main")

func ConnectDatabase() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(helper.GetENV("MONGODB_URI", "mongodb://downormal:downormal@localhost:27018/")))
	if err != nil {
		panic(err)
	}

	cdb := client.Database(helper.GetENV("DB_NAME", "db"))

	logger.Infof("connected to database: %v\n", cdb.Name())

	return cdb
}
