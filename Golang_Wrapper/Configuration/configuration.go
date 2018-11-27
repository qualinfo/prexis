package config

import (
	"log"

	"github.com/spf13/viper"
)

var configuration *Configuration

type Configuration struct {
	Prexis PrexisConfig
}

func GetConfig() *Configuration {
	return configuration
}

func Init() {

	viper.SetConfigName("config")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

}
