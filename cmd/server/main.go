// Package main is the entry point to the server. It reads configuration, sets up logging and error handling,
// handles signals from the OS, and starts and stops the server.
package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

func main() {
	os.Exit(start())
}

func start() int{
	logEnv := getStringOrDefault("LOG_ENV", "development")
	log, err := createLogger(logEnv)
	if err != nil{
		fmt.Println("Error setting up the logger", err)
		return 1
	}

	defer func() {
		_ = log.Sync()
	}()
	return 0
}


func createLogger(env string) (*zap.Logger, error){
	switch env {
	case  "production":
		return zap.NewProduction()
	case "development":
		return zap.NewDevelopment()
	default:
		return zap.NewNop(), nil
	}
}

func getStringOrDefault(name, defaultV string) string{
	v, ok := os.LookupEnv(name)
	if !ok{
		return defaultV
	}
	return v
}
