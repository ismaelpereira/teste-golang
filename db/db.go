package db

import (
	"context"
	"log"
	"teste-golang/types"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenDatabase(uri string) (*mongo.Client, error) {
	if uri == "" {
		log.Fatal("You must set an URI enviroment variable")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10+time.Second)
	defer cancel()

	m, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return m, err

}

func GetCollection(uri string) (*mongo.Collection, error) {
	client, err := OpenDatabase(uri)
	if err != nil {
		return nil, err
	}

	collection := client.Database("People_database").Collection("People")

	return collection, nil
}

func InsertInCollection(uri string, people types.People) error {
	collection, err := GetCollection(uri)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10+time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, people)

	spew.Dump(result)

	return err
}
