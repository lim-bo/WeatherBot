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
	currentWeatherURL = "https://api.openweathermap.org/data/2.5/weather"
	forecastURL       = "https://api.openweathermap.org/data/2.5/forecast"
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
				Timeout:   30 * time.Second,
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

	req, err := http.NewRequest(http.MethodGet, currentWeatherURL, nil)
	if err != nil {
		return nil, errors.New("owm error: " + err.Error())
	}
	q := req.URL.Query()
	q.Add("q", cityName)
	q.Add("appid", o.apiKey)
	q.Add("lang", "ru")
	req.URL.RawQuery = q.Encode()

	resp, err := o.cli.Do(req)
	if err != nil {
		return nil, errors.New("owm http error: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad request")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("owm reading error: " + err.Error())
	}
	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, errors.New("owm unmarshalling error: " + err.Error())
	}
	return &out, nil
}

var currentWeatherTmpl = "Локация: %s\nТекущая температура: %d°C, по ощущениям: %d°C\nСкорость ветра: %d м/c\nНаправление ветра: %d°"

func (o *OwmRepo) MakeCurrentWeatherCast(wc *entity.WeatherCast, cityName string) string {
	return fmt.Sprintf(currentWeatherTmpl, cityName, int16(wc.Main["temp"])-273,
		int16(wc.Main["feels_like"])-273,
		int16(wc.Wind["speed"]),
		int16(wc.Wind["deg"]),
	)
}

func (o *OwmRepo) Make3DayForecast(fc *entity.Forecast, cityName string) string {
	out := "Прогноз на 3 суток с интервалом в 3 часа (МСК: UTC+3)\n"
	timeStampTmpl := "-- %s Температура:%d°C, скорость ветра: %dм/c\n"
	var timeStr string
	for _, wc := range fc.List {
		timeStr = time.Unix(wc.Dt, 0).Format(time.DateTime)
		out += fmt.Sprintf(timeStampTmpl, timeStr, int16(wc.Main["temp"])-273, int16(wc.Wind["speed"]))
	}
	return out
}

func (o *OwmRepo) Get3DayForecast(cityName string) (*entity.Forecast, error) {
	req, err := http.NewRequest(http.MethodGet, forecastURL, nil)
	if err != nil {
		return nil, errors.New("owm error: " + err.Error())
	}
	q := req.URL.Query()
	q.Add("q", cityName)
	q.Add("appid", o.apiKey)
	q.Add("lang", "ru")
	// cnt is a count of timestamps, 8 units == 1 day
	q.Add("cnt", "24")
	req.URL.RawQuery = q.Encode()

	resp, err := o.cli.Do(req)
	if err != nil {
		return nil, errors.New("owm http error: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad request")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("owm reading error: " + err.Error())
	}
	var fc entity.Forecast
	err = json.Unmarshal(body, &fc)
	if err != nil {
		return nil, errors.New("owm unmarshalling error: " + err.Error())
	}
	return &fc, nil
}
