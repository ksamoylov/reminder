package config

import (
	pkg "reminder/pkg/env"
)

const (
	DefaultDbPort     = 5432
	DefaultServerPort = 8000
	TrueValue         = 1
)

type Config struct {
	DbConfig
	HttpPort  int
	DebugMode bool
}

func New() *Config {
	return &Config{
		DbConfig:  *newDbConfig(),
		HttpPort:  pkg.GetIntEnv("SERVER_PORT", DefaultServerPort),
		DebugMode: pkg.GetIntEnv("DEBUG_MODE", TrueValue) == TrueValue,
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
		Url:  pkg.GetEnv("POSTGRES_URL", ""),
		Host: pkg.GetEnv("POSTGRES_HOST", ""),
		Port: pkg.GetIntEnv("POSTGRES_PORT", DefaultDbPort),
		User: pkg.GetEnv("POSTGRES_USER", ""),
		Pass: pkg.GetEnv("POSTGRES_PASSWORD", ""),
		Name: pkg.GetEnv("POSTGRES_DB", ""),
	}
}
