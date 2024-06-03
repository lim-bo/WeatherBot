package bot

import (
	"fmt"
	"os"
	"weatherbot/logger"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	tgbot.BotAPI
}

// Constructor
func NewBot() (*Bot, error) {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		logger.LogFatalError(fmt.Errorf(`invalid bot wtoken`))
	}

	newBot, err := tgbot.NewBotAPI(botToken)
	logger.LogFatalError(err)
	return &Bot{*newBot}, nil
}

// Run-method which will be called from main
func (b *Bot) Serve() {
	upd, err := b.GetUpdatesChan(tgbot.NewUpdate(0))
	logger.LogFatalError(err)
	// TO-DO remove placeholder
	for _ = range upd {

	}
}
