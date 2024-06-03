package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
	"weatherbot/entity"
	"weatherbot/logger"
)

var requestTemplate = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&lang=ru"

type WeatherRepo interface {
	GetCurrentWeather(cityName string) (entity.WeatherCast, error)
}

// WeatherRepo implementation for OperWeatherMap api

type OwmRepo struct {
	cli    *http.Client
	apiKey string
}

func New() *OwmRepo {
	var owm OwmRepo
	key, ok := os.LookupEnv("WEATHER_API_KEY")
	if !ok {
		logger.LogFatalError(errors.New("cannot get apiKey"))
	}
	owm.apiKey = key
	cl := &http.Client{
		Timeout: time.Second * 20,
	}
	owm.cli = cl
	return &owm
}

func (o *OwmRepo) GetCurrentWeather(cityName string) (entity.WeatherCast, error) {
	var out entity.WeatherCast
	url, err := url.Parse(fmt.Sprintf(requestTemplate, cityName, o.apiKey))
	if err != nil {
		return out, errors.New("weather repo: " + err.Error())
	}

	req := http.Request{
		URL: url,
	}
	resp, err := o.cli.Do(&req)
	if err != nil {
		resp.Body.Close()
		return out, errors.New("weather repo: request: " + err.Error())
	}
	defer resp.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&out)
	if err != nil {
		return out, errors.New("weather repo: unmarshalling: " + err.Error())
	}
	if out.ResponseCode != http.StatusOK {
		return out, errors.New("weather repo: request: bad request")
	}
	return out, nil
}
