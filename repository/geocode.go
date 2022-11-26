package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/go-redis/redis/v8"
)

type GeocodeRepository interface {
	ReverseGeocode(ctx context.Context, lat, lng float64) (entity.ReverseGeocode, error)

	ReverseGeocodeCache(ctx context.Context, lat, lng float64) (entity.ReverseGeocode, error)
	SetReverseGeocodeCache(ctx context.Context, lat, lng float64, geocode entity.ReverseGeocode) error
}

type geocodeRepository struct {
	api         *entity.Config
	redisClient *redis.Client
}

func NewGeocodeRepository(api *entity.Config, redisClient *redis.Client) GeocodeRepository {
	return &geocodeRepository{api: api, redisClient: redisClient}
}

func (g *geocodeRepository) ReverseGeocodeCache(ctx context.Context, lat, lng float64) (entity.ReverseGeocode, error) {
	var data entity.ReverseGeocode

	key := fmt.Sprintf(geocodeKey, lat, lng)

	geocode, err := g.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return data, nil
	} else if err != nil {
		return data, err
	}

	if err = json.Unmarshal([]byte(geocode), &data); err != nil {
		return data, err
	}

	return data, nil
}

func (g *geocodeRepository) SetReverseGeocodeCache(ctx context.Context, lat, lng float64, geocode entity.ReverseGeocode) error {
	key := fmt.Sprintf(geocodeKey, lat, lng)

	geocodeByte, err := json.Marshal(geocode)
	if err != nil {
		return err
	}

	if err = g.redisClient.Set(ctx, key, string(geocodeByte), redisExpiration).Err(); err != nil {
		return err
	}

	return nil
}

func (g *geocodeRepository) ReverseGeocode(ctx context.Context, lat, lng float64) (entity.ReverseGeocode, error) {
	var data entity.ReverseGeocode

	body, err := helper.GetRequest(ctx, g.api.GeocodeRestApi+reverseGeocode(lat, lng))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
