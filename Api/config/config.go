package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT        string
	DELIVERY_SERVICE string
	TOKEN_KEY        string
	EMAIL_PASSWORD   string
	AUTH_SERVICE     string

	PostgresHost     string
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

	config.HTTP_PORT = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.DELIVERY_SERVICE = cast.ToString(getOrReturnDefaultValue("DELIVERY_SERVICE", ":8082"))
	config.AUTH_SERVICE = cast.ToString(getOrReturnDefaultValue("AUTH_SERVICE", ":8081"))
	config.EMAIL_PASSWORD = cast.ToString(getOrReturnDefaultValue("EMAIL_PASSWORD", "My-email-password"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "postgres-db"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1234"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "healthauth"))
	config.TOKEN_KEY = cast.ToString(getOrReturnDefaultValue("TOKEN_KEY", "my_secret_key"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
