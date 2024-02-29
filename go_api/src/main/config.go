package main

import "os"

type Config struct {
	Host string
	Port string

	// Redis Config
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

// _getEnv is a helper function to get the value of an environment variable or a default value
func _getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// LoadConfigFromEnviron loads the configuration from environment variables
func LoadConfigFromEnviron() Config {
	// Initialize the Config
	conf := Config{
		Host: _getEnv("HOST", "0.0.0.0"),
		Port: _getEnv("PORT", "8080"),

		// Redis Config
		RedisHost:     _getEnv("REDIS_HOST", "localhost"),
		RedisPort:     _getEnv("REDIS_PORT", "6379"),
		RedisPassword: _getEnv("REDIS_PASSWORD", ""),
		RedisDB:       0,
	}

	return conf
}
