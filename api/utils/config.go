package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresUser      string `mapstructure:"POSTGRES_USER"`
	PostgresPassword  string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb        string `mapstructure:"POSTGRES_DB"`
	PostgresHost      string `mapstructure:"POSTGRES_HOST"`
	PostgresPort      string `mapstructure:"POSTGRES_PORT"`
	ApiSecret         string `mapstructure:"API_SECRET"`
	TokenHourLifespan string `mapstructure:"TOKEN_HOUR_LIFESPAN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	// To get the value from the config file using key
	// viper package read .env

	if viper.ReadInConfig() != nil {
		fmt.Println("Failed to open config file")
	}

	err = viper.Unmarshal(&config)
	return
}
