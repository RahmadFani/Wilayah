package config

import (
	"os"
	"strings"
)

func New() {
	NewVars()
}

func GetSecret(path string, defaultEnv string, defaultValue string) string {
	data, err := os.ReadFile(path)
	if err == nil {
		// slog.Info("secret", "data", string(data))
		return strings.TrimSpace(string(data))
	}

	return GetEnv(defaultEnv, defaultValue)
}

func GetEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
