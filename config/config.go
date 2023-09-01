package config

import (
	pkg "reminder/pkg/env"
)

const (
	DefaultDbPort       = 5432
	DefaultServerPort   = 8000
	TrueValue           = 1
	DefaultPostgresPort = 6379
)

type Config struct {
	*DbConfig
	*RedisConfig
	HttpPort  int
	DebugMode bool
}

func New() *Config {
	return &Config{
		DbConfig:    newDbConfig(),
		RedisConfig: newRedisConfig(),
		HttpPort:    pkg.GetIntEnv("SERVER_PORT", DefaultServerPort),
		DebugMode:   pkg.GetIntEnv("DEBUG_MODE", TrueValue) == TrueValue,
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

type RedisConfig struct {
	Password string
	Db       int
	Host     string
	Port     int
}

func newRedisConfig() *RedisConfig {
	return &RedisConfig{
		Password: pkg.GetEnv("REDIS_PASSWORD", ""),
		Host:     pkg.GetEnv("REDIS_HOST", ""),
		Db:       pkg.GetIntEnv("REDIS_DB", 0),
		Port:     pkg.GetIntEnv("REDIS_PORT", 0),
	}
}
