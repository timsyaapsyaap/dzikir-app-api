package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/go-redis/redis/v8"
)

type HijriRepository interface {
	GregorianToHijri(ctx context.Context, date, month, year int) (entity.Hijri, error)

	GregorianToHijriCache(ctx context.Context, date, month, year int) (entity.Hijri, error)
	SetGregorianToHijriCache(ctx context.Context, date, month, year int, hijri entity.Hijri) error
}

type hijriRepository struct {
	api         *entity.Config
	redisClient *redis.Client
}

func NewHijriRepository(api *entity.Config, redisClient *redis.Client) HijriRepository {
	return &hijriRepository{
		api:         api,
		redisClient: redisClient,
	}
}

func (repository *hijriRepository) GregorianToHijri(ctx context.Context, date, month, year int) (entity.Hijri, error) {
	var (
		dataResponse entity.HijriReponseAPI
		data         entity.Hijri
	)

	body, err := helper.GetRequest(ctx, repository.api.HijriRestApi+gregorianToHijri(date, month, year))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	data = dataResponse.Data.Hijri
	data.Month.Id = monthIdHijri(data.Month.Number)

	return data, nil
}

func (repository *hijriRepository) GregorianToHijriCache(ctx context.Context, date, month, year int) (entity.Hijri, error) {
	var data entity.Hijri

	key := fmt.Sprintf(gregorianToHijriKey, date, month, year)

	hijri, err := repository.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return data, nil
	} else if err != nil {
		return data, err
	}

	if err = json.Unmarshal([]byte(hijri), &data); err != nil {
		return data, err
	}

	return data, nil
}

func (repository *hijriRepository) SetGregorianToHijriCache(ctx context.Context, date, month, year int, hijri entity.Hijri) error {
	key := fmt.Sprintf(gregorianToHijriKey, date, month, year)

	hijriByte, err := json.Marshal(hijri)
	if err != nil {
		return err
	}

	if err = repository.redisClient.Set(ctx, key, string(hijriByte), redisExpiration).Err(); err != nil {
		return err
	}

	return nil
}
