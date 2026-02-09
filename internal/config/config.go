// Package config обеспечивает загрузку и парсер конфигурации приложения.
package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config представляет конфигурацию приложения.
type Config struct {
	Name        string      `toml:"name"`
	Version     string      `toml:"version"`
	Description string      `toml:"description"`
	App         Logger      `toml:"app"`
	Bots        []BotConfig `toml:"bots"`
}

// Logger представляет уровень логированя приложения.
type Logger struct {
	LogLevel string `toml:"log_level"`
}

// BotConfig представляет конфигурацию управляемых ботов приложением.
type BotConfig struct {
	ID      string        `toml:"id"`
	Token   string        `toml:"token"`
	Enable  bool          `toml:"enable"`
	Storage StorageConfig `toml:"storage"`
}

// StorageConfig представляет конфигурацию подключения и монипуляцию приложения к базе данных.
type StorageConfig struct {
	Driver string `toml:"driver"`
	DSN    string `toml:"dsn"`
}

func MustLoad() *Config {
	var cfg Config
	if err := cleanenv.ReadConfig("./config.toml", &cfg); err != nil {
		log.Fatalf(
			"Could not read config file\nНе возможно прочитать файл конфигураци: %v",
			err,
		)
	}

	return &cfg
}
