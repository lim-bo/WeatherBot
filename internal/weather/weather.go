package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
	"weatherbot/entity"
	"weatherbot/logger"
)

var (
	requestURL = "https://api.openweathermap.org/data/2.5/weather"
)

// WeatherRepo implementation for OperWeatherMap api

type OwmRepo struct {
	cli    *http.Client
	apiKey string
	logger *logger.SLogger
}

func New(key string) *OwmRepo {
	sl := logger.NewSLogger()
	owm := OwmRepo{logger: sl}

	owm.apiKey = key
	cl := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 60 * time.Second,
		},
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

var currentWeatherTmpl = "Локация: %s\nТекущая температура: %d°C, по ощущениям: %d°C\nСкорость ветра: %d м/c\nНаправление ветра: %d°"

func (o *OwmRepo) MakeCurrentWeatherCast(wc *entity.WeatherCast, cityName string) string {
	return fmt.Sprintf(currentWeatherTmpl, cityName, int64(wc.Main["temp"])-273,
		int64(wc.Main["feels_like"])-273,
		int64(wc.Wind["speed"]),
		int64(wc.Wind["deg"]),
	)
}
