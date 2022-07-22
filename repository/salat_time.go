package repository

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	var dataResponse entity.SalatTimeCityFindRestAPIResponse
	var data []entity.SalatTimeCity

	body, err := helper.GetRequest(repository.api.SalatTimeRestApi + findCityAddress + city)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return nil, err
	}

	for _, v := range dataResponse.Data {
		idNumber, err := strconv.Atoi(v.ID)
		if err != nil {
			return data, err
		}

		data = append(data, entity.SalatTimeCity{
			ID:   idNumber,
			City: v.City,
		})
	}

	return data, nil
}

func (repository *salatTimeRepository) CityDetails(id string) (entity.SalatTimeCity, error) {
	var dataResponse entity.SalatTimeCityDetailsRestAPIResponse
	var data entity.SalatTimeCity

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
	var dataResponse []entity.SalatTimeCityRestAPI
	var data []entity.SalatTimeCity

	body, err := helper.GetRequest(repository.api.SalatTimeRestApi + allCities)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	for _, v := range dataResponse {
		idNumber, err := strconv.Atoi(v.ID)
		if err != nil {
			return data, err
		}

		data = append(data, entity.SalatTimeCity{
			ID:   idNumber,
			City: v.City,
		})
	}

	return data, nil
}

func (repository *salatTimeRepository) Schedule(cityId int, year int, month int, date int) (entity.SalatTime, error) {
	var dataResponse entity.SalatTimeRestAPIResponse
	var data entity.SalatTime

	body, err := helper.GetRequest(repository.api.SalatTimeRestApi + fmt.Sprintf("/v1/sholat/jadwal/%v/%v/%v/%v", cityId, year, month, date))
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
			FullDate: dataResponse.Data.Schedule.FullDate,
			Date:     dataResponse.Data.Schedule.Date,
			Imsak:    dataResponse.Data.Schedule.Imsak,
			Fajr:     dataResponse.Data.Schedule.Fajr,
			Rise:     dataResponse.Data.Schedule.Rise,
			Dhuhr:    dataResponse.Data.Schedule.Dhuhr,
			Asr:      dataResponse.Data.Schedule.Asr,
			Maghrib:  dataResponse.Data.Schedule.Maghrib,
			Isha:     dataResponse.Data.Schedule.Isha,
		},
	}

	return data, nil
}
