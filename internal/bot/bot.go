package bot

import (
	"context"
	"errors"
	"log/slog"
	"weatherbot/internal/weatherApi"
	"weatherbot/logger"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	WaitingToRespState int32 = 1
	DefaultState       int32 = 0
)

type Bot struct {
	api      tgbot.BotAPI
	Logger   *logger.SLogger
	WCClient weatherApi.WeatherCastServiceClient
}

// Constructor
func NewBot(token string, wcClient weatherApi.WeatherCastServiceClient) (*Bot, error) {
	sl := logger.New()
	newBot, err := tgbot.NewBotAPI(token)
	if err != nil {
		sl.Fatal(context.Background(), err)
	}
	return &Bot{
		api:      *newBot,
		Logger:   sl,
		WCClient: wcClient,
	}, nil
}

var keyBoard = tgbot.NewReplyKeyboard(
	tgbot.NewKeyboardButtonRow(
		tgbot.NewKeyboardButton("Текущая погода"),
		tgbot.NewKeyboardButton("Сменить город"),
		tgbot.NewKeyboardButton("Прогноз на 3 суток"),
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
		b.handleMessage(update)
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

func (b *Bot) handleMessage(update tgbot.Update) {
	ctx := context.Background()
	chatId := update.Message.Chat.ID
	// Checking if user with recieved chatid exist
	exist, err := b.WCClient.CheckUser(ctx, &weatherApi.UID{
		Value: chatId,
	})
	if err != nil {
		b.Logger.Error(ctx, err)
		return
	}

	// if not, creating new user in db
	if !exist.Value {

		_, err := b.WCClient.CreateUser(context.Background(), &weatherApi.UID{Value: chatId})
		if err != nil {
			b.Logger.Error(context.Background(), err)
			return
		}
		msg := tgbot.NewMessage(chatId, "Приветствую))")
		msg.ReplyMarkup = keyBoard
		if _, err := b.api.Send(msg); err != nil {
			b.Logger.Error(context.Background(), err)
		}
		b.Logger.LogWithGroupAtLevel(ctx, logger.LevelInfo,
			"new user registered",
			slog.String("user", update.Message.From.UserName))
	}
	// Recieving user to work with it
	user, err := b.WCClient.GetUser(context.Background(), &weatherApi.UID{Value: chatId})
	if err != nil {
		b.Logger.Error(context.Background(), err)
		return
	}
	// Changing prerfered city to weathercast
	if user.Status == WaitingToRespState {
		user.City = update.Message.Text
		user.Status = DefaultState
		_, err := b.api.Send(tgbot.NewMessage(chatId, "Принято"))
		if err != nil {
			b.Logger.Error(context.Background(), err)
			return
		}
		// refreshing user preference
		_, err = b.WCClient.SetUser(context.Background(), &weatherApi.User{
			Id:     user.Id,
			Status: user.Status,
			City:   user.City,
		})
		if err != nil {
			b.Logger.Error(context.Background(), err)
		}
		return
	}

	var command string
	switch update.Message.Text {
	case "Текущая погода":
		command = "current"
	case "Сменить город":
		command = "change"
	case "Прогноз на 3 суток":
		command = "3days"
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
		user.Status = WaitingToRespState
		_, err := b.WCClient.SetUser(context.Background(), &weatherApi.User{
			Id:     user.Id,
			City:   user.City,
			Status: user.Status,
		})
		if err != nil {
			b.Logger.Error(context.Background(), err)
			return
		}
		response := tgbot.NewMessage(chatId, "Укажите город, к которому хотели бы получать прогноз")
		if _, err := b.api.Send(response); err != nil {
			b.Logger.Error(context.Background(), err)
		}
	case "current":
		// Handling situation if user didn't choose city for cast
		if user.City == "" {
			b.SendWarn(chatId, "Город не выбран,\nдля получения прогноза укажите город")
			return
		}
		// Getting weatherCast from grpc service via client
		weatherCast, err := b.WCClient.GetCurrentWeather(context.Background(), &weatherApi.City{
			Name: user.City,
		})
		if err != nil {
			b.Logger.Error(context.Background(), err)
			return
		} else {
			b.Logger.LogWithGroupAtLevel(context.Background(),
				logger.LevelInfo,
				"weather request",
				slog.String("city", user.City),
				slog.String("user", update.Message.From.UserName),
			)
		}
		// Getting string representation of weathercast
		// from grpc service via client
		weatherCast.PrefCityName = user.City
		cast, _ := b.WCClient.MakeCurrentWeatherCast(context.Background(), weatherCast)
		_, err = b.api.Send(tgbot.NewMessage(chatId, cast.Text))
		if err != nil {
			b.Logger.Error(context.Background(), err)
		}
	case "3days":
		// Handling situation if user didn't choose city for cast
		if user.City == "" {
			b.SendWarn(chatId, "Город не выбран,\nдля получения прогноза укажите город")
			return
		}

		foreCast, err := b.WCClient.Get3DayForecast(context.Background(), &weatherApi.City{Name: user.City})
		if err != nil {
			b.Logger.Error(context.Background(), err)
			return
		} else {
			b.Logger.LogWithGroupAtLevel(context.Background(),
				logger.LevelInfo,
				"weather forecast request",
				slog.String("city", user.City),
				slog.String("user", update.Message.From.UserName),
			)
		}

		foreCast.PrefCityName = user.City
		cast, _ := b.WCClient.Make3DayForecast(context.Background(), foreCast)
		_, err = b.api.Send(tgbot.NewMessage(chatId, cast.Text))
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
}
