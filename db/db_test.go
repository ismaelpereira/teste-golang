package db

import (
	"log"
	"os"
	"teste-golang/types"
	"testing"

	"github.com/joho/godotenv"
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
	collection, err := GetCollection(uri)
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

	var people = types.People{ID: "1", Name: "John", Heigth: 180, Weigth: 80, Gender: "Male", IMC: ((80) / (180 / 10) * (180 / 10))}

	if err := InsertInCollection(uri, people); err != nil {
		t.Error("Cannot insert data", err)
	}
}
