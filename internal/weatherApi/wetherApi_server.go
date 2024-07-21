package weatherApi

import (
	"context"
	"weatherbot/entity"
	userdb "weatherbot/internal/userDB"
	"weatherbot/internal/weather"
)

// Repository for manage users' info
type UserManagerI interface {
	GetUserPreferences(int32) (string, error)
	SetUserPreference(int32, string) error
	CreateUserPreferences(int32, string) error
	GetUser(int32) (*userdb.User, error)
}

type WeatherApiServer struct {
	repo        *weather.OwmRepo
	userManager UserManagerI
	UnimplementedWeatherCastServiceServer
}

func NewWeatherApiSever(apikey string, dbCfg userdb.DBConfig) *WeatherApiServer {
	return &WeatherApiServer{
		repo: weather.New(apikey),
	}
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
