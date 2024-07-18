package main

import (
	"context"
	"log"
	"weatherbot/internal/bot"
	"weatherbot/internal/weatherApi"
	"weatherbot/logger"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
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

	b, err := bot.NewBot(v.GetString("BOT_TOKEN"))
	if err != nil {
		logger.NewSLogger().Fatal(context.Background(), err)
	}
	dial, err := grpc.NewClient(":8081", grpc.WithInsecure())
	if err != nil {
		b.Logger.Fatal(context.Background(), err)
	}
	cli := weatherApi.NewWeatherCastServiceClient(dial)
	b.WCClient = cli
	b.Serve()
}
