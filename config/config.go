package config

import (
	"os"
	"strconv"
)

const (
	DefaultPort = 5432
)

type Config struct {
	DbConfig
	TelegramConfig
}

func NewConfig() *Config {
	return &Config{
		DbConfig:       *newDbConfig(),
		TelegramConfig: *newTelegramConfig(),
	}
}

type DbConfig struct {
	Url  string
	Host string
	Port int
	User string
	Pass string
	Name string
}

func newDbConfig() *DbConfig {
	return &DbConfig{
		Url:  getEnv("POSTGRES_URL", ""),
		Host: getEnv("POSTGRES_HOST", ""),
		Port: getIntEnv("POSTGRES_PORT", DefaultPort),
		User: getEnv("POSTGRES_USER", ""),
		Pass: getEnv("POSTGRES_PASS", ""),
		Name: getEnv("POSTGRES_NAME", ""),
	}
}

type TelegramConfig struct {
	Token string
}

func newTelegramConfig() *TelegramConfig {
	return &TelegramConfig{
		Token: getEnv("TELEGRAM_BOT_TOKEN", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getIntEnv(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		val, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return defaultVal
		}

		return int(val)
	}

	return defaultVal
}
