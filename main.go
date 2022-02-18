package main

import (
	"fmt"
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
	fmt.Println(uri)
	client, err := db.OpenDatabase(uri)
	if err != nil {
		return err
	}

	spew.Dump(client)
	return nil
}
