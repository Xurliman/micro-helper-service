package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetDatabaseUser() string {
	return getEnvironmentValue("DB_USERNAME")
}

func GetDatabasePassword() string {
	return getEnvironmentValue("DB_PASSWORD")
}

func GetDatabaseName() string {
	return getEnvironmentValue("DB_NAME")
}

func GetAppPort() string {
	return getEnvironmentValue("APP_PORT")
}

func getEnvironmentValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return os.Getenv(key)
}
