package main

import (
	"log"
	"os"
	"teste-golang/db"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := godotenv.Load("local.env"); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	collection, err := db.OpenDatabase(uri)
	if err != nil {
		return err
	}

	spew.Dump(collection)
	return nil
}
