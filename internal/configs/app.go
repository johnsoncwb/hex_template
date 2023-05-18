package configs

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/newrelic"
	"os"
)

type AppConfig struct {
	AppName string
	Env     string
	Port    string

	NewRelicKey         string
	NewRelicApp         *newrelic.Application
	NewRelicTransaction *newrelic.Transaction
	NewRelicContext     context.Context
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

	cfg.NewRelicKey = os.Getenv("NEW_RELIC_LICENSE_KEY")

	return
}
