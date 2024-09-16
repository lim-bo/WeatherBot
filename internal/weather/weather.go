package weather

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
	"weatherbot/entity"
	"weatherbot/logger"

	"github.com/enescakir/emoji"
)

var (
	currentWeatherURL = "https://api.openweathermap.org/data/2.5/weather"
	forecastURL       = "https://api.openweathermap.org/data/2.5/forecast"
	ErrBadRequest     = errors.New("bad request")
)

// WeatherRepo implementation for OperWeatherMap api

type OwmRepo struct {
	cli    *http.Client
	apiKey string
	logger *logger.SLogger
}

func New(key string) *OwmRepo {
	sl := logger.New()
	owm := OwmRepo{logger: sl}
	owm.apiKey = key

	rootCAPool := x509.NewCertPool()
	rootCA, err := ioutil.ReadFile("./certs/client-cert/cert.pem")
	if err != nil {
		sl.Fatal(context.Background(), err)
	}
	rootCAPool.AppendCertsFromPEM(rootCA)

	cl := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 60 * time.Second,
			TLSClientConfig:     &tls.Config{RootCAs: rootCAPool},
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
		return nil, ErrBadRequest
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("owm reading error: " + err.Error())
	}
	err = out.UnmarshalJSON(body)
	if err != nil {
		return nil, errors.New("owm unmarshalling error: " + err.Error())
	}
	return &out, nil
}

var currentWeatherTmpl = "Город: %s\nТекущая температура: %d°C, по ощущениям: %d°C\nСкорость ветра: %d м/c\nНаправление ветра: %d°\n"
var freezzingWeatherMessage = "На улице реальная зима" + emoji.ColdFace.String() +
	"\nСоветую закутаться как капуста)" + emoji.Gloves.String() + emoji.Scarf.String()
var coldWeatherMessage = "Однако прохладно" + emoji.ConfusedFace.String() + "\nБудет не лишним накинуть легкую куртку" + emoji.Coat.String() +
	"\nНа всякий случай следует захватить зонтик" + emoji.WinkingFace.String() + emoji.Umbrella.String()
var warmWeatherMessage = "Тёпленько..." + emoji.SmilingFace.String() + "\nНо для пляжных шорт пока рановато" + emoji.WinkingFace.String()
var hotWeatherMessage = "Печка" + emoji.HotFace.String() + "\nНастало время для легкой одежды и шоколадного мороженного" + emoji.TShirt.String() + emoji.Dress.String() + emoji.IceCream.String()

func (o *OwmRepo) MakeCurrentWeatherCast(wc *entity.WeatherCast, cityName string) string {
	main := fmt.Sprintf(currentWeatherTmpl, cityName, int16(wc.Main["temp"])-273,
		int16(wc.Main["feels_like"])-273,
		int16(wc.Wind["speed"]),
		int16(wc.Wind["deg"]),
	)
	switch {
	case wc.Main["temp"] < -5.0:
		main += freezzingWeatherMessage
	case wc.Main["temp"] > -5.0 && wc.Main["temp"] < 10.0:
		main += coldWeatherMessage
	case wc.Main["temp"] > 10.0 && wc.Main["temp"] < 20.0:
		main += warmWeatherMessage
	case wc.Main["temp"] > 20.0:
		main += hotWeatherMessage
	}
	return main
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
		return nil, ErrBadRequest
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil || len(body) == 0 {
		return nil, errors.New("owm reading error: " + err.Error())
	}
	var fc entity.Forecast
	err = fc.UnmarshalJSON(body)
	if err != nil {
		return nil, errors.New("owm unmarshalling error: " + err.Error())
	}
	return &fc, nil
}
