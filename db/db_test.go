package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDb(t *testing.T) {
	if err := godotenv.Load("local.env"); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	collection, err := OpenDatabase(uri)
	if err != nil {
		t.Error("Fail to open database")
	}
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		t.Error("Fail to open database")
	}

	var parsedCollection []People
	if err = cursor.All(context.TODO(), &parsedCollection); err != nil {
		t.Error("Fail to open database")
	}

}
