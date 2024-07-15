package bot

import (
	"context"
	"fmt"
	"log/slog"
	"weatherbot/entity"
	"weatherbot/logger"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

type WCRepoI interface {
	GetCurrentWeather(string) (*entity.WeatherCast, error)
}

type Bot struct {
	api    tgbot.BotAPI
	Logger *logger.SLogger
	WCRepo WCRepoI
}

var currentWeatherTmpl = "Локация: %s\nТекущая температура: %d°C, по ощущениям: %d°C\nСкорость ветра: %d м/c\nНаправление ветра: %d°"

// Constructor
func NewBot(token string, wcRepo WCRepoI) (*Bot, error) {
	sl := logger.NewSLogger()
	newBot, err := tgbot.NewBotAPI(token)
	if err != nil {
		sl.Fatal(context.Background(), err)
	}
	return &Bot{
		api:    *newBot,
		Logger: sl,
		WCRepo: wcRepo,
	}, nil
}

// Run-method
func (b *Bot) Serve() {
	upd, err := b.api.GetUpdatesChan(tgbot.NewUpdate(0))
	if err != nil {
		b.Logger.Fatal(context.Background(), err)
	}
	b.Logger.Info(context.Background(), "Bot started gettings updates")
	for update := range upd {
		go func(update tgbot.Update) {
			switch update.Message.Command() {
			case "current":
				weatherCast, err := b.WCRepo.GetCurrentWeather("Norilsk")
				if err != nil {
					b.Logger.Error(context.Background(), err)
					return
				} else {
					b.Logger.LogWithGroupAtLevel(context.Background(),
						logger.LevelInfo,
						"weather request",
						slog.String("city", update.Message.Text),
						slog.String("user", update.Message.From.UserName),
					)
				}
				reponseString := fmt.Sprintf(currentWeatherTmpl, "Norilsk", int64(weatherCast.Main["temp"])-273,
					int64(weatherCast.Main["feels_like"])-273,
					int64(weatherCast.Wind["speed"]),
					int64(weatherCast.Wind["deg"]),
				)
				response := tgbot.NewMessage(update.Message.Chat.ID, reponseString)
				_, err = b.api.Send(response)
				if err != nil {
					b.Logger.Error(context.Background(), err)
				}
			default:
				_, err = b.api.Send(tgbot.NewMessage(update.Message.Chat.ID, "Not supported yet"))
				if err != nil {
					b.Logger.Error(context.Background(), err)
				}
			}

		}(update)
	}
}
