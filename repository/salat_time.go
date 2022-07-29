package repository

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/helper"
)

type SalatTimeRepository interface {
	FindCity(city string) ([]entity.SalatTimeCity, error)
	CityDetails(id string) (entity.SalatTimeCity, error)
	AllCities() ([]entity.SalatTimeCity, error)
	Schedule(cityId int, year int, month int, date int) (entity.SalatTime, error)
}

type salatTimeRepository struct {
	api *entity.Config
}

func NewSalatRepository(api *entity.Config) SalatTimeRepository {
	return &salatTimeRepository{
		api: api,
	}
}

func (repository *salatTimeRepository) FindCity(city string) ([]entity.SalatTimeCity, error) {
	var (
		dataResponse entity.SalatTimeCityFindRestAPIResponse
		data         []entity.SalatTimeCity
		wg           sync.WaitGroup
	)

	body, err := helper.GetRequest(repository.api.SalatTimeRestApi + findCityAddress + city)
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
			defer wg.Done()

			idNumber, err := strconv.Atoi(value.ID)
			if err != nil {
				panic(err)
			}

			data = append(data, entity.SalatTimeCity{
				ID:   idNumber,
				City: value.City,
			})
		}(v)
	}
	wg.Wait()

	return data, nil
}

func (repository *salatTimeRepository) CityDetails(id string) (entity.SalatTimeCity, error) {
	var (
		dataResponse entity.SalatTimeCityDetailsRestAPIResponse
		data         entity.SalatTimeCity
	)

	body, err := helper.GetRequest(repository.api.SalatTimeRestApi + cityDetails + id)
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

func (repository *salatTimeRepository) AllCities() ([]entity.SalatTimeCity, error) {
	var (
		dataResponse []entity.SalatTimeCityRestAPI
		data         []entity.SalatTimeCity
		wg           sync.WaitGroup
	)

	body, err := helper.GetRequest(repository.api.SalatTimeRestApi + allCities)
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
			defer wg.Done()

			idNumber, err := strconv.Atoi(value.ID)
			if err != nil {
				panic(err)
			}

			data = append(data, entity.SalatTimeCity{
				ID:   idNumber,
				City: value.City,
			})
		}(v)

	}
	wg.Wait()

	return data, nil
}

func (repository *salatTimeRepository) Schedule(cityId int, year int, month int, date int) (entity.SalatTime, error) {
	var (
		dataResponse entity.SalatTimeRestAPIResponse
		data         entity.SalatTime
	)

	body, err := helper.GetRequest(repository.api.SalatTimeRestApi + schedule(cityId, year, month, date))
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
