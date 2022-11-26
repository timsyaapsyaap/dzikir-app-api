package service

import (
	"context"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
)

type GeocodeService interface {
	ReverseGeocode(ctx context.Context, lat, lng float64) (entity.ReverseGeocode, error)
}

type geocodeService struct {
	geocodeRepository repository.GeocodeRepository
}

func NewGeocodeService(geocodeRepository repository.GeocodeRepository) GeocodeService {
	return &geocodeService{geocodeRepository: geocodeRepository}
}

func (g *geocodeService) ReverseGeocode(ctx context.Context, lat, lng float64) (entity.ReverseGeocode, error) {
	var geocode entity.ReverseGeocode

	geocode, err := g.geocodeRepository.ReverseGeocodeCache(ctx, lat, lng)
	if err != nil {
		return geocode, err
	}

	if geocode.City == "" && geocode.Locality == "" {
		geocode, err = g.geocodeRepository.ReverseGeocode(ctx, lat, lng)
		if err != nil {
			return geocode, err
		}

		if err = g.geocodeRepository.SetReverseGeocodeCache(ctx, lat, lng, geocode); err != nil {
			return geocode, err
		}
	}

	return geocode, nil
}
