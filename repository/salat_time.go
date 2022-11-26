package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/go-redis/redis/v8"
)

type SalatTimeRepository interface {
	FindCity(ctx context.Context, city string) ([]entity.SalatTimeCity, error)
	CityDetails(ctx context.Context, id string) (entity.SalatTimeCity, error)
	AllCities(ctx context.Context) ([]entity.SalatTimeCity, error)
	Schedule(ctx context.Context, cityId, year, month, date int) (entity.SalatTime, error)

	AllCitiesCache(ctx context.Context) ([]entity.SalatTimeCity, error)
	SetAllCitiesCache(ctx context.Context, cities []entity.SalatTimeCity) error
	CityDetailsCache(ctx context.Context, id string) (entity.SalatTimeCity, error)
	SetCityDetailsCache(ctx context.Context, id string, city entity.SalatTimeCity) error
	FindCitiesCache(ctx context.Context, city string) ([]entity.SalatTimeCity, error)
	SetFindCitiesCache(ctx context.Context, city string, cities []entity.SalatTimeCity) error
	ScheduleCache(ctx context.Context, cityId, year, month, date int) (entity.SalatTime, error)
	SetScheduleCache(ctx context.Context, cityId, year, month, date int, salatTime entity.SalatTime) error
}

type salatTimeRepository struct {
	api         *entity.Config
	redisClient *redis.Client
}

func NewSalatRepository(api *entity.Config, redisClient *redis.Client) SalatTimeRepository {
	return &salatTimeRepository{
		api:         api,
		redisClient: redisClient,
	}
}

func (repository *salatTimeRepository) FindCity(ctx context.Context, city string) ([]entity.SalatTimeCity, error) {
	var (
		dataResponse entity.SalatTimeCityFindRestAPIResponse
		data         []entity.SalatTimeCity
		wg           sync.WaitGroup
		mutex        sync.Mutex
	)

	body, err := helper.GetRequest(ctx, repository.api.SalatTimeRestApi+findCityAddress+city)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return nil, err
	}

	wg.Add(len(dataResponse.Data))
	for _, v := range dataResponse.Data {
		go func(value entity.SalatTimeCityRestAPI) {
			defer mutex.Unlock()
			defer wg.Done()

			idNumber, err := strconv.Atoi(value.ID)
			if err != nil {
				panic(err)
			}

			mutex.Lock()
			data = append(data, entity.SalatTimeCity{
				ID:   idNumber,
				City: value.City,
			})
		}(v)
	}
	wg.Wait()

	return data, nil
}

func (repository *salatTimeRepository) CityDetails(ctx context.Context, id string) (entity.SalatTimeCity, error) {
	var (
		dataResponse entity.SalatTimeCityDetailsRestAPIResponse
		data         entity.SalatTimeCity
	)

	body, err := helper.GetRequest(ctx, repository.api.SalatTimeRestApi+cityDetails+id)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return data, err
	}

	data = entity.SalatTimeCity{
		ID:   idNumber,
		City: dataResponse.Data.City,
	}

	return data, nil
}

func (repository *salatTimeRepository) AllCities(ctx context.Context) ([]entity.SalatTimeCity, error) {
	var (
		dataResponse []entity.SalatTimeCityRestAPI
		data         []entity.SalatTimeCity
		wg           sync.WaitGroup
		mutex        sync.Mutex
	)

	body, err := helper.GetRequest(ctx, repository.api.SalatTimeRestApi+allCities)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	wg.Add(len(dataResponse))
	for _, v := range dataResponse {
		go func(value entity.SalatTimeCityRestAPI) {
			defer mutex.Unlock()
			defer wg.Done()

			idNumber, err := strconv.Atoi(value.ID)
			if err != nil {
				panic(err)
			}

			mutex.Lock()
			data = append(data, entity.SalatTimeCity{
				ID:   idNumber,
				City: value.City,
			})
		}(v)

	}
	wg.Wait()

	return data, nil
}

func (repository *salatTimeRepository) Schedule(ctx context.Context, cityId, year, month, date int) (entity.SalatTime, error) {
	var (
		dataResponse entity.SalatTimeRestAPIResponse
		data         entity.SalatTime
	)

	body, err := helper.GetRequest(ctx, repository.api.SalatTimeRestApi+schedule(cityId, year, month, date))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	id, err := strconv.Atoi(dataResponse.Data.ID)
	if err != nil {
		return data, err
	}

	data = entity.SalatTime{
		ID:       id,
		City:     dataResponse.Data.City,
		Province: dataResponse.Data.Province,
		Coordinate: entity.SalatTimeCoordinate{
			Lat:       dataResponse.Data.Coordinate.Lat,
			Lng:       dataResponse.Data.Coordinate.Lng,
			Latitude:  dataResponse.Data.Coordinate.Latitude,
			Longitude: dataResponse.Data.Coordinate.Longitude,
		},
		Schedule: entity.SalatTimeSchedule{
			Imsak:   dataResponse.Data.Schedule.Imsak,
			Fajr:    dataResponse.Data.Schedule.Fajr,
			Rise:    dataResponse.Data.Schedule.Rise,
			Duha:    dataResponse.Data.Schedule.Duha,
			Dhuhr:   dataResponse.Data.Schedule.Dhuhr,
			Asr:     dataResponse.Data.Schedule.Asr,
			Maghrib: dataResponse.Data.Schedule.Maghrib,
			Isha:    dataResponse.Data.Schedule.Isha,
		},
	}

	return data, nil
}

func (repository *salatTimeRepository) AllCitiesCache(ctx context.Context) ([]entity.SalatTimeCity, error) {
	var data []entity.SalatTimeCity

	cities, err := repository.redisClient.Get(ctx, allCitiesKey).Result()
	if err == redis.Nil {
		return data, nil
	} else if err != nil {
		return data, err
	}

	if err = json.Unmarshal([]byte(cities), &data); err != nil {
		return data, err
	}

	return data, nil
}

func (repository *salatTimeRepository) SetAllCitiesCache(ctx context.Context, cities []entity.SalatTimeCity) error {
	citiesByte, err := json.Marshal(cities)
	if err != nil {
		return err
	}

	if err = repository.redisClient.Set(ctx, allCitiesKey, string(citiesByte), redisExpiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repository *salatTimeRepository) CityDetailsCache(ctx context.Context, id string) (entity.SalatTimeCity, error) {
	var data entity.SalatTimeCity

	key := fmt.Sprintf(findCityDetailsKey, id)

	city, err := repository.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return data, nil
	} else if err != nil {
		return data, err
	}

	if err = json.Unmarshal([]byte(city), &data); err != nil {
		return data, err
	}

	return data, nil
}

func (repository *salatTimeRepository) SetCityDetailsCache(ctx context.Context, id string, city entity.SalatTimeCity) error {
	key := fmt.Sprintf(findCityDetailsKey, id)

	cityByte, err := json.Marshal(city)
	if err != nil {
		return err
	}

	if err = repository.redisClient.Set(ctx, key, string(cityByte), redisExpiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repository *salatTimeRepository) FindCitiesCache(ctx context.Context, city string) ([]entity.SalatTimeCity, error) {
	var data []entity.SalatTimeCity

	key := fmt.Sprintf(findCityDetailsKey, city)

	cities, err := repository.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return data, nil
	} else if err != nil {
		return data, err
	}

	if err = json.Unmarshal([]byte(cities), &data); err != nil {
		return data, err
	}

	return data, nil
}

func (repository *salatTimeRepository) SetFindCitiesCache(ctx context.Context, city string, cities []entity.SalatTimeCity) error {
	key := fmt.Sprintf(findCityDetailsKey, city)

	citiesByte, err := json.Marshal(cities)
	if err != nil {
		return err
	}

	if err = repository.redisClient.Set(ctx, key, string(citiesByte), redisExpiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repository *salatTimeRepository) ScheduleCache(ctx context.Context, cityId, year, month, date int) (entity.SalatTime, error) {
	var data entity.SalatTime

	key := fmt.Sprintf(scheduleKey, cityId, year, month, date)

	salatTime, err := repository.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return data, nil
	} else if err != nil {
		return data, err
	}

	if err = json.Unmarshal([]byte(salatTime), &data); err != nil {
		return data, err
	}

	return data, nil
}

func (repository *salatTimeRepository) SetScheduleCache(ctx context.Context, cityId, year, month, date int, salatTime entity.SalatTime) error {
	key := fmt.Sprintf(scheduleKey, cityId, year, month, date)

	salatTimeByte, err := json.Marshal(salatTime)
	if err != nil {
		return err
	}

	if err = repository.redisClient.Set(ctx, key, string(salatTimeByte), redisExpiration).Err(); err != nil {
		return err
	}

	return nil
}
