package service

import (
	"context"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
)

type SalatTimeService interface {
	FindCity(ctx context.Context, city string) ([]entity.SalatTimeCity, error)
	CityDetails(ctx context.Context, id string) (entity.SalatTimeCity, error)
	AllCities(ctx context.Context) ([]entity.SalatTimeCity, error)
	Schedule(ctx context.Context, cityId int, year int, month int, date int) (entity.SalatTime, error)
}

type salatTimeService struct {
	salatTimeRepository repository.SalatTimeRepository
}

func NewSalatTimeService(salatTimeRepository repository.SalatTimeRepository) SalatTimeService {
	return &salatTimeService{
		salatTimeRepository: salatTimeRepository,
	}
}

func (service *salatTimeService) FindCity(ctx context.Context, city string) ([]entity.SalatTimeCity, error) {
	var cities []entity.SalatTimeCity

	cities, err := service.salatTimeRepository.FindCitiesCache(ctx, city)
	if err != nil {
		return cities, err
	}

	if len(cities) == 0 {
		cities, err = service.salatTimeRepository.FindCity(ctx, city)
		if err != nil {
			return cities, err
		}

		if err = service.salatTimeRepository.SetFindCitiesCache(ctx, city, cities); err != nil {
			return cities, err
		}
	}

	return cities, nil
}

func (service *salatTimeService) CityDetails(ctx context.Context, id string) (entity.SalatTimeCity, error) {
	var city entity.SalatTimeCity

	city, err := service.salatTimeRepository.CityDetailsCache(ctx, id)
	if err != nil {
		return city, err
	}

	if city.ID == 0 {
		city, err = service.salatTimeRepository.CityDetails(ctx, id)
		if err != nil {
			return city, err
		}

		if err = service.salatTimeRepository.SetCityDetailsCache(ctx, id, city); err != nil {
			return city, err
		}
	}

	return city, nil
}

func (service *salatTimeService) AllCities(ctx context.Context) ([]entity.SalatTimeCity, error) {
	var cities []entity.SalatTimeCity

	cities, err := service.salatTimeRepository.AllCitiesCache(ctx)
	if err != nil {
		return cities, err
	}

	if len(cities) == 0 {
		cities, err = service.salatTimeRepository.AllCities(ctx)
		if err != nil {
			return cities, err
		}

		if err = service.salatTimeRepository.SetAllCitiesCache(ctx, cities); err != nil {
			return cities, err
		}
	}

	return cities, nil
}

func (service *salatTimeService) Schedule(ctx context.Context, cityId, year, month, date int) (entity.SalatTime, error) {
	var salatTime entity.SalatTime

	salatTime, err := service.salatTimeRepository.ScheduleCache(ctx, cityId, year, month, date)
	if err != nil {
		return salatTime, err
	}

	if salatTime.ID == 0 {
		salatTime, err = service.salatTimeRepository.Schedule(ctx, cityId, year, month, date)
		if err != nil {
			return salatTime, err
		}

		if err = service.salatTimeRepository.SetScheduleCache(ctx, cityId, year, month, date, salatTime); err != nil {
			return salatTime, err
		}
	}

	return salatTime, nil
}
