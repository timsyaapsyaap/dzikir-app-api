package config

import (
	"log"
	"os"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

// Setup environment variables
func SetupEnvironment() *entity.Config {
	env := os.Getenv("ENV")
	if env != "production" {
		// Get the config from .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// Set environment variables
	salatTimeRestApi := os.Getenv("SALAT_TIME_REST_API")
	quranRestApi := os.Getenv("QURAN_REST_API")
	hijriRestApi := os.Getenv("HIJRI_REST_API")
	geocodeRestApi := os.Getenv("GEOCODE_REST_API")
	redisUrl := os.Getenv("REDIS_URL")

	return &entity.Config{
		SalatTimeRestApi: salatTimeRestApi,
		QuranRestApi:     quranRestApi,
		HijriRestApi:     hijriRestApi,
		GeocodeRestApi:   geocodeRestApi,
		RedisURL:         redisUrl,
	}
}

// Redis
func NewRedisConn(config *entity.Config) *redis.Client {
	opt, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)

	return rdb
}
