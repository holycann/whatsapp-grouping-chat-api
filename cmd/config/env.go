package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PublicHost   string
	Port         string
	DBAddress    string
	MaxOpenConns int64
	MaxIdleConns int64
	MaxIdleTime  int64
}

var Env = initConfigProduction()

func initConfigProduction() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "https://wise-sheena-holycan-e9914987.koyeb.app"),
		Port:       getEnv("PORT", "8080"),
		DBAddress: fmt.Sprintf(
			"postgres://%s:%s@%s/%s",
			getEnv("DB_USER", "holycan"),
			getEnv("DB_PASSWORD", "Vjh38sroTQql"),
			getEnv("DB_HOST", "ep-mute-forest-a1sttr12.ap-southeast-1.pg.koyeb.app"),
			getEnv("DB_NAME", "koyebdb"),
		),
		MaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  getEnvAsInt("DB_MAX_IDLE_TIME", 5),
	}
}

func initConfigStaging() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBAddress: fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=%s",
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "ramaa212!"),
			getEnv("DB_HOST", "localhost:5432"),
			getEnv("DB_NAME", "postgres"),
			getEnv("SSL_MODE", "disable"),
		),
		MaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 30),
		MaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 30),
		MaxIdleTime:  getEnvAsInt("DB_MAX_IDLE_TIME", 5),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
