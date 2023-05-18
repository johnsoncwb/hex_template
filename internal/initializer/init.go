package initializer

import (
	"github.com/johnsoncwb/hex_template/internal/configs"
)

func LoadEnv() {
	err := configs.GlobalConfig.LoadVariables()
	if err != nil {
		panic(err)
	}
}

func StartApp() {
	LoadEnv()

}
