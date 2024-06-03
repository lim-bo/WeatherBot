package main

import (
	"weatherbot/logger"

	"github.com/joho/godotenv"
)

func main() {
	InitEnv()

	// b, err := bot.NewBot()
	// logger.LogFatalError(err)

}

func InitEnv() {
	// load values from .env file
	if err := godotenv.Load("../.env"); err != nil {
		logger.LogFatalError(err)
	}
}
