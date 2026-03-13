package storage

import (
	"context"

	"github.com/iamthefamous/goscrapper/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Collection, error) {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)

	if err != nil {
		return nil, err
	}

	colection := client.Database("scraper").Collection("articles")
	return colection, nil
}

func SaveArticles(collection *mongo.Collection, articles []models.Article) error {
	var docs []interface{}
	for _, article := range articles {
		docs = append(docs, article)
	}

	_, err := collection.InsertMany(context.TODO(), docs)
	return err
}
