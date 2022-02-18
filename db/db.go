package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type People struct {
	ID     string `json:"_id" bson:"_id"`
	Name   string
	Heigth int
	Weigth int
	IMC    float64
	Gender string
}

func OpenDatabase(uri string) (*mongo.Collection, error) {

	if uri == "" {
		log.Fatal("You must set an URI enviroment variable")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10+time.Second)
	defer cancel()

	m, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := m.Disconnect(ctx); err != nil {
			panic(err)
		}

	}()

	log.Println("Database Running")

	collection := m.Database("People_database").Collection("People")

	return collection, nil

}
