package repository

import (
	"encoding/json"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/helper"
)

type HijriRepository interface {
	GregorianToHijri(date int, month int, year int) (entity.Hijri, error)
}

type hijriRepository struct {
	api *entity.Config
}

func NewHijriRepository(api *entity.Config) HijriRepository {
	return &hijriRepository{
		api: api,
	}
}

func (repository *hijriRepository) GregorianToHijri(date int, month int, year int) (entity.Hijri, error) {
	var (
		dataResponse entity.HijriReponseAPI
		data         entity.Hijri
	)

	body, err := helper.GetRequest(repository.api.HijriRestApi + gregorianToHijri(date, month, year))
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
