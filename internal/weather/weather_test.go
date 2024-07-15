package weather_test

import (
	"errors"
	"log"
	"testing"
	"weatherbot/internal/weather"

	"github.com/spf13/viper"
)

type testCase struct {
	CityName string
	Expect   error
}

var v *viper.Viper

func TestMain(m *testing.M) {
	v = viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("secret")
	v.AddConfigPath("../../configs")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("file not found", err)
		} else {

			log.Fatal("cfg reading error", err)
		}
	}
}

func TestCurrentWeather(t *testing.T) {
	testCases := []testCase{
		{
			"Москва",
			nil,
		},
		{
			"Moscow",
			nil,
		},
		{
			"asdajdkajwasd",
			errors.New("bad request"),
		},
	}
	testRepo := weather.New(v.GetString("WEATHER_API_KEY"))
	for _, cs := range testCases {
		cast, err := testRepo.GetCurrentWeather(cs.CityName)
		if err != nil && err.Error() != cs.Expect.Error() {
			t.Fatalf("cityName: %s\n error: %s\n expected: %s\n", cs.CityName, err.Error(), cs.Expect.Error())
		}
		t.Log(cast)
	}
}
