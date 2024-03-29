package pkg

import (
	"os"
	"strconv"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func GetIntEnv(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		val, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return defaultVal
		}

		return int(val)
	}

	return defaultVal
}
