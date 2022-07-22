package service

import (
	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
)

type SalatTimeService interface {
	FindCity(city string) ([]entity.SalatTimeCity, error)
	CityDetails(id string) (entity.SalatTimeCity, error)
	AllCities() ([]entity.SalatTimeCity, error)
	Schedule(cityId int, year int, month int, date int) (entity.SalatTime, error)
}

type salatTimeService struct {
	salatTimeRepository repository.SalatTimeRepository
}

func NewSalatTimeService(salatTimeRepository repository.SalatTimeRepository) SalatTimeService {
	return &salatTimeService{
		salatTimeRepository: salatTimeRepository,
	}
}

func (service *salatTimeService) FindCity(city string) ([]entity.SalatTimeCity, error) {
	return service.salatTimeRepository.FindCity(city)
}

func (service *salatTimeService) CityDetails(id string) (entity.SalatTimeCity, error) {
	return service.salatTimeRepository.CityDetails(id)
}

func (service *salatTimeService) AllCities() ([]entity.SalatTimeCity, error) {
	return service.salatTimeRepository.AllCities()
}

func (service *salatTimeService) Schedule(cityId int, year int, month int, date int) (entity.SalatTime, error) {
	return service.salatTimeRepository.Schedule(cityId, year, month, date)
}
