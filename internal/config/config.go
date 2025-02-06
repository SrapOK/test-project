package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HttpServer  HttpServer  `mapstructure:"http_server"`
	SongDetails SongDetails `mapstructure:"song_details"`
	DB          DB
	LogLevel    string `mapstructure:"log_level"`
}

type HttpServer struct {
	Address     string
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
}

type DB struct {
	DSN string
}

type SongDetails struct {
	Url string
}

func Load(pathToConfig string) (*Config, error) {
	config := new(Config)
	viper.SetConfigFile(pathToConfig)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("файл конфигурации не найден: %s", err.Error())
		}

		return nil, fmt.Errorf("не удаётся загрузить файл с конфигурацией: %s", err.Error())
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("не удаётся десериализовать конфигурацию: %s", err.Error())
	}

	viper.AutomaticEnv()

	return config, nil
}
