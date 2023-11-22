package util

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost        string `mapstructure:"DB_HOST"`
	DBName        string `mapstructure:"DB_NAME"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBPort        string `mapstructure:"DB_PORT"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
}

var Conf Config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Error().Err(err).Msg("viper err read config")
	}
	if err = viper.Unmarshal(&Conf); err != nil {
		log.Error().Err(err).Msg("viper err unmarshaling config")
	}

}
