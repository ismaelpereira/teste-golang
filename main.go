package main

import (
	"log"
	"teste-golang/api"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	api.StartAPI()
	return nil
}
