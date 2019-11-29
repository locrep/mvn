package db

import (
	"context"
	"github.com/locrep/go/config"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Client struct {
	artifact *mongo.Collection
}

func Connect(mongoUrl string) Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + mongoUrl))
	if err != nil {
		logger.WithFields(CouldntConnectMongoServer(err)).Error(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), config.Conf.DBConnectionTimeout*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logger.WithFields(CouldntConnectMongoDB(err)).Error(err.Error())
	}

	return Client{artifact: client.Database("maven").Collection("artifact")}
}

func (c Client) GetAll() {
	ctx, _ := context.WithTimeout(context.Background(), config.Conf.DBReadTimeout*time.Second)
	cur, err := c.artifact.Find(ctx, nil)
	if err != nil {
		logger.WithFields(CouldntReadArtifactsFromMongoDB(err)).Error(err.Error())
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			logger.WithFields(CouldntDecodeArtifact(err)).Error(err.Error())
		}
	}

	if err := cur.Err(); err != nil {
		logger.WithFields(GotErrorFromArtifactCursor(err)).Error(err.Error())
	}
}

func (c Client) FindArtifact() {
	ctx, _ := context.WithTimeout(context.Background(), config.Conf.DBReadTimeout*time.Second)
	cur, err := c.artifact.Find(ctx, nil)
	if err != nil {
		logger.WithFields(CouldntReadArtifactsFromMongoDB(err)).Error(err.Error())
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			logger.WithFields(CouldntDecodeArtifact(err)).Error(err.Error())
		}
	}

	if err := cur.Err(); err != nil {
		logger.WithFields(GotErrorFromArtifactCursor(err)).Error(err.Error())
	}
}
