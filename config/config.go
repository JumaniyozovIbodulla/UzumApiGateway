package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const ()

// Config ...
type Config struct {
	Environment   string // develop, staging, production
	RedisHost     string
	RedisPort     int
	RedisPassword string

	OrderServiceHost string
	OrderServicePort string

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", ""))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", ""))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ""))
	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", ""))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))
	c.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", ""))

	c.OrderServiceHost = cast.ToString(getOrReturnDefault("Order_SERVICE_HOST", "localhost"))
	c.OrderServicePort = cast.ToString(getOrReturnDefault("Order_GRPC_PORT", "8080"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
