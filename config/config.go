package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func GetDatabaseConnection() string {
	return getEnvironmentValue("DB_CONNECTION")
}

func GetDatabaseUser() string {
	return getEnvironmentValue("DB_USERNAME")
}

func GetDatabasePassword() string {
	return getEnvironmentValue("DB_PASSWORD")
}

func GetDatabaseName() string {
	return getEnvironmentValue("DB_NAME")
}

func GetDatabaseHost() string {
	return getEnvironmentValue("DB_HOST")
}

func GetDatabasePort() string {
	return getEnvironmentValue("DB_PORT")
}

func GetAppPort() string {
	return getEnvironmentValue("APP_PORT")
}

func getEnvironmentValue(key string) string {
	err := godotenv.Load(dir(".env"))
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return os.Getenv(key)
}

func dir(envFile string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			panic(fmt.Errorf("go.mod not found"))
		}
		currentDir = parent
	}

	return filepath.Join(currentDir, envFile)
}
