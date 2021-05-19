package mongodb

import (
	"context"

	"github.com/fazarrahman/blogbe/lib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New() (*mongo.Database, error) {
	username := lib.GetEnv("DB_USERNAME")
	password := lib.GetEnv("DB_PASSWORD")
	host := lib.GetEnv("DB_HOST")
	port := lib.GetEnv("DB_PORT")
	dbname := lib.GetEnv("DB_NAME")

	uri := host + ":" + port
	if username != "" && password != "" {
		uri = username + ":" + password + "@" + uri
	}

	clientOptions := options.Client().ApplyURI("mongodb://" + uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return client.Database(dbname), nil
}
