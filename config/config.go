package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string
	// ...
}

func Init() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cnf := &Config{}
	if err := viper.Unmarshal(cnf); err != nil {
		return nil, err
	}
	return cnf, nil
}
