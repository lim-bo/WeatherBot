package weatherApi

import (
	"context"
	"weatherbot/entity"
	userdb "weatherbot/internal/userDB"
	"weatherbot/internal/weather"
	"weatherbot/logger"
)

// Repository for manage users' info
type UserManagerI interface {
	GetUser(int64) (*userdb.User, error)
	SetUser(userdb.User) error
	CheckUserExist(int64) (bool, error)
	CreateUser(int64) error
}

type WeatherRepo interface {
	GetCurrentWeather(string) (*entity.WeatherCast, error)
	MakeCurrentWeatherCast(*entity.WeatherCast, string) string
	Get3DayForecast(string) (*entity.Forecast, error)
	Make3DayForecast(*entity.Forecast, string) string
}

type WeatherApiServer struct {
	repo        WeatherRepo
	userManager UserManagerI
	lg          *logger.SLogger

	UnimplementedWeatherCastServiceServer
}

func NewWeatherApiServer(apikey string, dbCfg userdb.DBConfig) *WeatherApiServer {
	return &WeatherApiServer{
		repo:        weather.New(apikey),
		userManager: userdb.NewUserDB(dbCfg),
		lg:          logger.NewSLogger(),
	}
}

func (srv *WeatherApiServer) GetCurrentWeather(ctx context.Context, city *City) (*WeatherCast, error) {
	wc, err := srv.repo.GetCurrentWeather(city.Name)
	if err != nil {
		srv.lg.Error(context.Background(), err)
		return nil, err
	}
	return &WeatherCast{
		Main:       wc.Main,
		StatusCode: int32(wc.ResponseCode),
		Wind:       wc.Wind,
	}, nil
}

func (srv *WeatherApiServer) Get3DayForecast(ctx context.Context, city *City) (*Forecast, error) {
	fc, err := srv.repo.Get3DayForecast(city.Name)
	if err != nil {
		srv.lg.Error(context.Background(), err)
		return nil, err
	}
	list := make([]*WeatherCast, 0)
	// Need to "cast" repo-layer entities to proto types
	for _, wc := range fc.List {
		list = append(list, &WeatherCast{
			Main:     wc.Main,
			Wind:     wc.Wind,
			Datetime: wc.Dt,
		})
	}
	return &Forecast{
		StatusCode:   int32(fc.ResponseCode),
		List:         list,
		PrefCityName: city.Name,
	}, nil
}

func (srv *WeatherApiServer) Make3DayForecast(ctx context.Context, fc *Forecast) (*Cast, error) {
	list := make([]*entity.ForecastUnit, 0)
	for _, wc := range fc.List {
		list = append(list, &entity.ForecastUnit{
			Main: wc.Main,
			Wind: wc.Wind,
			// Adding 3 hours in sec to get UTC+3
			Dt: wc.Datetime + 10800,
		})
	}
	out := srv.repo.Make3DayForecast(&entity.Forecast{
		List: list,
	}, fc.PrefCityName)
	return &Cast{Text: out}, nil
}

func (srv *WeatherApiServer) MakeCurrentWeatherCast(ctx context.Context, wc *WeatherCast) (*Cast, error) {
	out := srv.repo.MakeCurrentWeatherCast(&entity.WeatherCast{
		Main:         wc.Main,
		Wind:         wc.Wind,
		ResponseCode: int(wc.StatusCode),
	}, wc.PrefCityName)
	return &Cast{Text: out}, nil
}

func (srv *WeatherApiServer) GetUser(ctx context.Context, id *UID) (*User, error) {
	u, err := srv.userManager.GetUser(id.Value)
	if err != nil {
		srv.lg.Error(context.Background(), err)
		return nil, err
	}
	return &User{
		Id:     u.Id,
		City:   u.City,
		Status: u.Status,
	}, nil
}

func (srv *WeatherApiServer) CheckUser(ctx context.Context, id *UID) (*IsExist, error) {
	exist, err := srv.userManager.CheckUserExist(id.Value)
	if err != nil {
		srv.lg.Error(context.Background(), err)
		return nil, err
	}
	return &IsExist{
		Value: exist,
	}, nil
}

// *Error is a placeholder because of
// how rpc works, get it with _, err := ...
// TO-DO: fix it
func (srv *WeatherApiServer) SetUser(ctx context.Context, u *User) (*Error, error) {
	err := srv.userManager.SetUser(userdb.User{
		Id:     u.Id,
		City:   u.City,
		Status: u.Status,
	})
	if err != nil {
		srv.lg.Error(context.Background(), err)
		return nil, err
	}
	return nil, nil
}

func (srv *WeatherApiServer) CreateUser(ctx context.Context, id *UID) (*Error, error) {
	err := srv.userManager.CreateUser(id.Value)
	if err != nil {
		srv.lg.Error(context.Background(), err)
		return nil, err
	}
	return nil, nil
}
