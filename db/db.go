package db

import (
	"context"
	"fmt"
	"log"
	"teste-golang/types"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func FindCollection(uri string) (*mongo.Collection, error) {
	client, err := OpenDatabase(uri)
	if err != nil {
		return nil, err
	}

	collection := client.Database("People_database").Collection("People")

	return collection, nil
}

func FindOneInCollection(uri string, ID primitive.ObjectID) (primitive.M, error) {
	client, err := OpenDatabase(uri)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10+time.Second)
	defer cancel()

	collection := client.Database("People_database").Collection("People")

	var people bson.M

	err = collection.FindOne(ctx, bson.D{{"_id", ID}}).Decode(&people)

	if err != nil {
		return nil, err
	}
	return people, nil
}

func CheckIfNameExists(uri string, name string) (bool, error) {
	client, err := OpenDatabase(uri)
	if err != nil {
		return true, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10+time.Second)
	defer cancel()

	collection := client.Database("People_database").Collection("People")

	var people bson.M

	err = collection.FindOne(ctx, bson.D{{"name", name}}).Decode(&people)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return true, err
	}
	if people != nil {
		return true, err
	}

	return false, nil
}

func InsertInCollection(uri string, people types.People) error {
	collection, err := FindCollection(uri)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10+time.Second)
	defer cancel()

	exists, err := CheckIfNameExists(uri, people.Name)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("ERROR! Name already exists")
	}

	result, err := collection.InsertOne(ctx, people)
	if err != nil {
		return err
	}

	fmt.Println("Cadastro da pessoa ", result)

	return err
}

func UpdatePeople(uri string, id primitive.ObjectID, people types.People) error {
	collection, err := FindCollection(uri)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10+time.Second)
	defer cancel()

	filter := bson.D{{"_id", people.ID}}

	spew.Dump(filter)

	result, err := collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: people}})
	if err != nil {
		return err
	}

	fmt.Println("Edição da pessoa ", result)
	return nil
}

func DeletePeople(uri string, id primitive.ObjectID) error {
	collection, err := FindCollection(uri)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10+time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	fmt.Println("Exclusão da pessoa ", result)
	return nil
}
