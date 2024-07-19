package bot

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"weatherbot/internal/weatherApi"
	"weatherbot/logger"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Type for represent user preferences about
// cuty for cast
type userState struct {
	isAwaitingToResp bool
	chosenCity       string
}

type Bot struct {
	api        tgbot.BotAPI
	Logger     *logger.SLogger
	WCClient   weatherApi.WeatherCastServiceClient
	mu         sync.RWMutex
	userStates map[int64]*userState
}

// Constructor
func NewBot(token string) (*Bot, error) {
	sl := logger.NewSLogger()
	newBot, err := tgbot.NewBotAPI(token)
	if err != nil {
		sl.Fatal(context.Background(), err)
	}
	return &Bot{
		api:        *newBot,
		Logger:     sl,
		mu:         sync.RWMutex{},
		userStates: make(map[int64]*userState, 0),
	}, nil
}

var keyBoard = tgbot.NewReplyKeyboard(
	tgbot.NewKeyboardButtonRow(
		tgbot.NewKeyboardButton("Текущая погода"),
		tgbot.NewKeyboardButton("Сменить город"),
	),
)

// Run-method
func (b *Bot) Serve() {
	// Check if object hasn't init-ed client
	if b.WCClient == nil {
		b.Logger.Fatal(context.Background(), errors.New("bot need a rpc client for start"))
	}

	// Getting channel for recieving updated
	upd, err := b.api.GetUpdatesChan(tgbot.NewUpdate(0))
	if err != nil {
		b.Logger.Fatal(context.Background(), err)
	}

	// There is a need to clear all messages, which
	// were sent while bot was off, to avoid ddos or another sh*t
	upd.Clear()

	b.Logger.Info(context.Background(), "Bot started gettings updates")
	for update := range upd {
		if update.Message == nil { // ignore non-Message updates
			continue
		}
		go func(update tgbot.Update) {
			chatId := update.Message.Chat.ID
			b.mu.Lock()
			if _, ok := b.userStates[chatId]; !ok {
				b.userStates[chatId] = &userState{}
				msg := tgbot.NewMessage(chatId, "Приветствую))")
				msg.ReplyMarkup = keyBoard
				if _, err := b.api.Send(msg); err != nil {
					b.Logger.Error(context.Background(), err)
				}
			}
			b.mu.Unlock()
			// Changing prerfered city to weathercast
			b.mu.Lock()
			if !update.Message.IsCommand() && b.userStates[chatId].isAwaitingToResp {
				b.userStates[chatId].chosenCity = update.Message.Text
				b.userStates[chatId].isAwaitingToResp = false
				b.mu.Unlock()
				_, err := b.api.Send(tgbot.NewMessage(chatId, "Принято"))
				if err != nil {
					b.Logger.Error(context.Background(), err)
				}
				return
			}
			b.mu.Unlock()
			var command string
			switch update.Message.Text {
			case "Текущая погода":
				command = "current"
			case "Сменить город":
				command = "change"
			}
			if command == "" {
				b.SendWarn(chatId, "Выберите команду из предложенных")
				return
			}

			// Handling commands
			switch command {
			case "change":
				// Getting bot ready to recieve name of the cuty
				// which weather user want to get
				b.mu.Lock()
				b.userStates[chatId].isAwaitingToResp = true
				b.mu.Unlock()
				response := tgbot.NewMessage(chatId, "Укажите город, к которому хотели бы получать прогноз")
				if _, err := b.api.Send(response); err != nil {
					b.Logger.Error(context.Background(), err)
				}

			case "current":
				b.mu.RLock()
				// Handling situation if user didn't choose city for cast
				if b.userStates[chatId].chosenCity == "" {
					b.SendWarn(chatId, "Город не выбран,\nдля получения прогноза укажите город")
					b.mu.RUnlock()
					return
				}
				// Getting weatherCast from grpc service via client
				weatherCast, err := b.WCClient.GetCurrentWeather(context.Background(), &weatherApi.City{
					Name: b.userStates[chatId].chosenCity,
				})
				b.mu.RUnlock()
				if err != nil {
					b.Logger.Error(context.Background(), err)
					return
				} else {
					b.Logger.LogWithGroupAtLevel(context.Background(),
						logger.LevelInfo,
						"weather request",
						slog.String("city", b.userStates[chatId].chosenCity),
						slog.String("user", update.Message.From.UserName),
					)
				}
				// Getting string representation of weathercast
				// from grpc service via client
				weatherCast.PrefCityName = b.userStates[chatId].chosenCity
				cast, err := b.WCClient.MakeCurrentWeatherCast(context.Background(), weatherCast)
				if err != nil {
					b.Logger.Error(context.Background(), err)
					return
				}
				response := tgbot.NewMessage(chatId, cast.Text)
				_, err = b.api.Send(response)
				if err != nil {
					b.Logger.Error(context.Background(), err)
				}
			default:
				b.Logger.LogWithGroupAtLevel(context.Background(),
					logger.LevelTrace,
					"unsupported message",
					slog.String("msg", update.Message.Text),
				)
				b.SendWarn(chatId, "Выберите команду из предложенных")
			}

		}(update)
	}
}

// Send warning message if incoming request
// is not supported
func (b *Bot) SendWarn(chatId int64, text string) {
	_, err := b.api.Send(tgbot.NewMessage(chatId, text))
	if err != nil {
		b.Logger.Error(context.Background(), err)
	}
}
