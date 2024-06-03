package entity

type WeatherCast struct {
	Coord        map[string]float64 `json:"coord"`
	Main         map[string]float64 `json:"main"`
	Wind         map[string]float64 `json:"wind"`
	ResponseCode int                `json:"cod"`
}
