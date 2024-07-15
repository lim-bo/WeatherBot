package main

import (
	"context"
	"log"
	"weatherbot/internal/bot"
	"weatherbot/internal/weather"
	"weatherbot/logger"

	"github.com/spf13/viper"
)

var v *viper.Viper

func init() {
	// loading config
	v = viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("secret")
	v.AddConfigPath("./configs")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("file not found", err)
		} else {

			log.Fatal("cfg reading error", err)
		}
	}
}

func main() {

	b, err := bot.NewBot(v.GetString("BOT_TOKEN"), weather.New(v.GetString("WEATHER_API_KEY")))
	if err != nil {
		logger.NewSLogger().Fatal(context.Background(), err)
	}
	b.Serve()
}
