package weather_test

import (
	"errors"
	"testing"
	"weatherbot/internal/weather"

	"github.com/joho/godotenv"
)

type testCase struct {
	CityName string
	Expect   error
}

func TestCurrentWeather(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(errors.New(".env loading fail"))
	}
	testCases := []testCase{
		{
			"Москва",
			nil,
		},
		{
			"asdajdkajwasd",
			errors.New("bad request"),
		},
	}
	testRepo := weather.New()
	for _, cs := range testCases {
		cast, err := testRepo.GetCurrentWeather(cs.CityName)
		if err != nil && err.Error() != cs.Expect.Error() {
			t.Fatalf("cityName: %s\n error: %s\n expected: %s\n", cs.CityName, err.Error(), cs.Expect.Error())
		}
		t.Log(cast)
	}
}
