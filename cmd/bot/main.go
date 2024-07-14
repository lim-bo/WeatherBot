package main

import (
	"context"
	"weatherbot/logger"

	"github.com/joho/godotenv"
)

func init() {
	// load values from .env file
	if err := godotenv.Load("../.env"); err != nil {
		logger.NewSLogger().Fatal(context.Background(), err)
	}
}

func main() {

	// b, err := bot.NewBot()
	// logger.NewSLogger().Fatal(context.Background(), err)

}
