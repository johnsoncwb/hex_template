package initializer

import (
	"context"
	"fmt"
	"github.com/johnsoncwb/hex_template/internal/configs"
	"github.com/newrelic/go-agent/v3/newrelic"
	log "github.com/sirupsen/logrus"
	"time"
)

func LoadEnv() {
	err := configs.GlobalConfig.LoadVariables()
	if err != nil {
		panic(err)
	}
}

func StartApp() {
	LoadEnv()
	ctx := context.Background()
	StartNewRelic(ctx)

}

func StartNewRelic(ctx context.Context) {

	var err error

	configs.GlobalConfig.NewRelicApp, err = newrelic.NewApplication(
		newrelic.ConfigAppName(fmt.Sprintf("%s-%s", configs.GlobalConfig.AppName, configs.GlobalConfig.Env)),
		newrelic.ConfigLicense(configs.GlobalConfig.NewRelicKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Erro ao iniciar a NewRelic")
	}

	err = configs.GlobalConfig.NewRelicApp.WaitForConnection(5 * time.Second)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Erro ao conectar com a NewRelic")
	}

}
