package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	AppName string
	Env     string
	Port    string
}

var GlobalConfig AppConfig

func (cfg *AppConfig) LoadVariables() (err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	cfg.AppName = os.Getenv("APP_NAME")
	cfg.Env = os.Getenv("APP_ENV")
	cfg.Port = os.Getenv("PORT")

	return
}
