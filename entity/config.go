package entity

// Config represents configuration of application.
type Config struct {
	SalatTimeRestApi string `env:"SALAT_TIME_REST_API"`
	QuranRestApi     string `env:"QURAN_REST_API"`
	RedisURL         string `env:"REDIS_URL"`
}
