package weather_test

import (
	"errors"
	"testing"
	"weatherbot/pkg/weather"
)

type testCase struct {
	CityName string
	Expect   error
}

func TestCurrentWeather(t *testing.T) {
	testCases := []testCase{
		{
			"Москва",
			nil,
		},
		{
			"asdajdkajwasd",
			errors.New("weather repo: request: bad request"),
		},
	}
	testRepo := weather.New()
	for _, cs := range testCases {
		_, err := testRepo.GetCurrentWeather(cs.CityName)
		if err != nil && err.Error() != cs.Expect.Error() {
			t.Fatalf("cityName: %s\n error: %s\n expected: %s\n", cs.CityName, err.Error(), cs.Expect.Error())
		}
	}
}
