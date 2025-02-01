package config

import (
	modelconfig "cs-ticketing/internal/models/config"
	"encoding/json"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (appConfig modelconfig.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev.json")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&appConfig)
	PrintConfig(appConfig)

	return
}

func PrintConfig(c modelconfig.Config) {
	data, _ := json.MarshalIndent(c, "", "\t")
	log.Println(string(data))
}
