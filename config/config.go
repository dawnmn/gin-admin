package config

import (
	"github.com/spf13/viper"
	"log"
)

type Server struct {
	App App
}

type App struct {
	DebugMode    bool
	CookieSecret string
	LogPath      string
	Host         string
	Port         string
}

func InitConfig(config *Server) {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := v.Unmarshal(config); err != nil {
		log.Fatal(err)
	}
}
