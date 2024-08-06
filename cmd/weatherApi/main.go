package main

import (
	"log"
	"net"
	userdb "weatherbot/internal/userDB"
	"weatherbot/internal/weatherApi"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var v *viper.Viper

func init() {
	// loading config
	v = viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("secret")
	v.AddConfigPath("./configs")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("file not found", err)
		} else {

			log.Fatal("cfg reading error", err)
		}
	}
}

func main() {
	dbCfg := userdb.DBConfig{
		User:   v.GetString("DB_USER"),
		Pass:   v.GetString("DB_PASSWORD"),
		DBName: v.GetString("DB_NAME"),
		Host:   "postgres_container",
		Port:   "5432",
	}
	srv := weatherApi.NewWeatherApiServer(v.GetString("WEATHER_API_KEY"), dbCfg)
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	weatherApi.RegisterWeatherCastServiceServer(s, srv)
	log.Println("weatherCast service listening on :8081")
	log.Fatal(s.Serve(lis))
}
