package bot

import (
	"context"
	"errors"
	"os"
	"weatherbot/logger"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	api    tgbot.BotAPI
	logger *logger.SLogger
}

// Constructor
func NewBot() (*Bot, error) {
	sl := logger.NewSLogger()
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		sl.Fatal(context.Background(), errors.New(`invalid bot wtoken`))
	}

	newBot, err := tgbot.NewBotAPI(botToken)
	if err != nil {
		sl.Fatal(context.Background(), err)
	}
	return &Bot{
		api:    *newBot,
		logger: sl,
	}, nil
}

// Run-method which will be called from main
func (b *Bot) Serve() {
	upd, err := b.api.GetUpdatesChan(tgbot.NewUpdate(0))
	if err != nil {
		b.logger.Fatal(context.Background(), err)
	}
	// TO-DO remove placeholder
	for _ = range upd {

	}
}
