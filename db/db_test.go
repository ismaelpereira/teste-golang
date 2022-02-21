package db

import (
	"log"
	"os"
	"teste-golang/types"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestOpenDb(t *testing.T) {
	if err := godotenv.Load("../local.env"); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	client, err := OpenDatabase(uri)
	if err != nil {
		t.Error("Fail to open database")
	}
	if client == nil {
		t.Error("Client cannot open")
	}

}

func TestGetCollection(t *testing.T) {
	if err := godotenv.Load("../local.env"); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	collection, err := FindCollection(uri)
	if err != nil {
		t.Error("Fail to get collection", err)
	}

	if collection == nil {
		t.Error("Collection does not exists", err)
	}
}

func TestInsertion(t *testing.T) {
	if err := godotenv.Load("../local.env"); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	var people = types.People{ID: primitive.NewObjectID(), Name: "Carlos Silva", Heigth: 180, Weigth: 80, Gender: "Male", IMC: ((80) / (180 / 10) * (180 / 10))}

	if err := InsertInCollection(uri, people); err == nil {
		t.Error("Cannot insert data", err)
	}
}

func TestUpdate(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("621015f8c710919c4e49270b")
	if err != nil {
		t.Error("Error, ID doesn't exist", err)
	}
	if err := godotenv.Load("../local.env"); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	var people = types.People{ID: id, Name: "Fernanda", Heigth: 155, Weigth: 45, Gender: "Female", IMC: (45 / (155 / 10) * (155 / 10))}

	err = UpdatePeople(uri, id, people)
	if err != nil {
		t.Error("Cannot update user", err)
	}

}

func TestDelete(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("620ed0dc2f2971556973785c")
	if err != nil {
		t.Error("Error, ID doesn't exist", err)
	}
	if err := godotenv.Load("../local.env"); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	err = DeletePeople(uri, id)
	if err != nil {
		t.Error("Cannot delete user", err)
	}
}
