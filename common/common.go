package common

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func LoadUri(path string) (string, error) {
	if err := godotenv.Load(path); err != nil {
		return "", fmt.Errorf("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	return uri, nil
}

func Context() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10+time.Second)

	return ctx, cancel
}
