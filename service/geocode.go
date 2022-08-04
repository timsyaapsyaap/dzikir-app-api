package service

import (
	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
)

type GeocodeService interface {
	ReverseGeocode(lat float64, lng float64) (entity.ReverseGeocode, error)
}

type geocodeService struct {
	geocodeRepository repository.GeocodeRepository
}

func NewGeocodeService(geocodeRepository repository.GeocodeRepository) GeocodeService {
	return &geocodeService{geocodeRepository: geocodeRepository}
}

func (g *geocodeService) ReverseGeocode(lat float64, lng float64) (entity.ReverseGeocode, error) {
	return g.geocodeRepository.ReverseGeocode(lat, lng)
}
