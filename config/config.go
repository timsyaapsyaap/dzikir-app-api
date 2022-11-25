package config

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/gomodule/redigo/redis"
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
func NewPool(config entity.Config) *redis.Pool {
	url := config.RedisURL

	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(url, redis.DialTLSSkipVerify(true), redis.DialTLSConfig(&tls.Config{InsecureSkipVerify: true}))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
