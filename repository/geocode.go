package repository

import (
	"encoding/json"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/helper"
)

type GeocodeRepository interface {
	ReverseGeocode(lat float64, lng float64) (entity.ReverseGeocode, error)
}

type geocodeRepository struct {
	api *entity.Config
}

func NewGeocodeRepository(api *entity.Config) GeocodeRepository {
	return &geocodeRepository{api: api}
}

func (g *geocodeRepository) ReverseGeocode(lat float64, lng float64) (entity.ReverseGeocode, error) {
	var data entity.ReverseGeocode

	body, err := helper.GetRequest(g.api.GeocodeRestApi + reverseGeocode(lat, lng))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
