package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT      string
	TOKEN_KEY      string
	REDIS_HOST     string
	REDIS_PORT     int
	EMAIL_PASSWORD string

	POSTGRES_HOST    string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8081"))
	config.REDIS_HOST = cast.ToString(getOrReturnDefaultValue("REDIS_HOST", "redis"))
	config.REDIS_PORT = cast.ToInt(getOrReturnDefaultValue("REDIS_PORT", 6379))

	config.POSTGRES_HOST = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "postgres-db"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1234"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "delivery"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
