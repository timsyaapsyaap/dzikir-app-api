package config

import (
	"os"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/subosito/gotenv"
)

// Setup environment variables
func SetupEnvironment() *entity.Config {
	err := gotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Set environment variables
	salatTimeRestApi := os.Getenv("SALAT_TIME_REST_API")
	quranRestApi := os.Getenv("QURAN_REST_API")

	return &entity.Config{
		SalatTimeRestApi: salatTimeRestApi,
		QuranRestApi:     quranRestApi,
	}
}
