package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
	AppPort string
}

var Env *Environment

func getEnv(key string, required bool) string {
	value, ok := os.LookupEnv(key)

	if !ok && required {
		log.Fatalf("Missing or invalid environment key: '%s'", key)
	}

	return value
}

func LoadEnvironment() {
	if Env == nil {
		Env = new(Environment)
	}

	Env.DBHost = getEnv("DB_HOST", true)
	Env.DBPort = getEnv("DB_PORT", true)
	Env.DBUser = getEnv("DB_USERNAME", true)
	Env.DBPass = getEnv("DB_PASSWORD", true)
	Env.DBName = getEnv("DB_NAME", true)
	Env.AppPort = getEnv("APP_PORT", true)
}

func LoadEnvironmentFile(file string) {
	if err := godotenv.Load(file); err != nil {
		fmt.Printf("Error on load environment file: %s", file)
	}

	LoadEnvironment()
}
