package weatherApi

import (
	"context"
	"weatherbot/entity"
	"weatherbot/internal/weather"
)

type WeatherApiServer struct {
	repo *weather.OwmRepo
	UnimplementedWeatherCastServiceServer
}

func (srv *WeatherApiServer) GetCurrentWeather(ctx context.Context, city *City) (*WeatherCast, error) {
	wc, err := srv.repo.GetCurrentWeather(city.Name)
	if err != nil {
		return nil, err
	}
	return &WeatherCast{
		Main:       wc.Main,
		StatusCode: int32(wc.ResponseCode),
		Wind:       wc.Wind,
	}, nil
}

func (srv *WeatherApiServer) MakeCurrentWeatherCast(ctx context.Context, wc *WeatherCast) (*Cast, error) {
	out := srv.repo.MakeCurrentWeatherCast(&entity.WeatherCast{
		Main:         wc.Main,
		Wind:         wc.Wind,
		ResponseCode: int(wc.StatusCode),
	}, wc.PrefCityName)
	return &Cast{Text: out}, nil
}

func NewWeatherApiSever(apikey string) *WeatherApiServer {
	return &WeatherApiServer{
		repo: weather.New(apikey),
	}
}
