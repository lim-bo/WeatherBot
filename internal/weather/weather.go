package weather

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
	"weatherbot/entity"
	"weatherbot/logger"
)

var (
	requestURL = "https://api.openweathermap.org/data/2.5/weather"
)

type WeatherRepo interface {
	GetCurrentWeather(cityName string) (entity.WeatherCast, error)
}

// WeatherRepo implementation for OperWeatherMap api

type OwmRepo struct {
	cli    *http.Client
	apiKey string
	logger *logger.SLogger
}

func New() *OwmRepo {
	sl := logger.NewSLogger()
	owm := OwmRepo{logger: sl}
	key, ok := os.LookupEnv("WEATHER_API_KEY")
	if !ok {
		owm.logger.Fatal(context.Background(), errors.New("cannot get apiKey"))
	}
	owm.apiKey = key
	cl := &http.Client{
		Timeout: time.Second * 20,
	}
	owm.cli = cl
	return &owm
}

func (o *OwmRepo) GetCurrentWeather(cityName string) (*entity.WeatherCast, error) {
	out := entity.WeatherCast{}

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("q", cityName)
	q.Add("appid", o.apiKey)
	q.Add("lang", "ru")
	req.URL.RawQuery = q.Encode()

	resp, err := o.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad request")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
