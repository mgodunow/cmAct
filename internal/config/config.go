package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbPass			string `mapstructure:"DB_PASS"`
	SecretPhrase	string `mapstructure:"SECRET_PHRASE"`
	Salt 			string `mapstructure:"SALT"`
}

func loadConfig() (config Config, err error) {
	viper.AddConfigPath("./../../")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}