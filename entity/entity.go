package entity

//easyjson:json
type WeatherCast struct {
	Coord        map[string]float64 `json:"coord"`
	Main         map[string]float64 `json:"main"`
	Wind         map[string]float64 `json:"wind"`
	ResponseCode int                `json:"cod"`
}

type WeatherCastError struct {
	ResponseCode int    `json:"cod"`
	Message      string `json:"message"`
	Params       []string
}

//easyjson:json
type ForecastUnit struct {
	Dt   int64              `json:"dt"`
	Main map[string]float64 `json:"main"`
	Wind map[string]float64 `json:"wind"`
}

//easyjson:json
type Forecast struct {
	ResponseCode string         `json:"cod"`
	List         []ForecastUnit `json:"list"`
}
