package service

import (
	"context"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
)

type HijriService interface {
	GregorianToHijri(ctx context.Context, date, month, year int) (entity.Hijri, error)
}

type hijriService struct {
	hijriRepository repository.HijriRepository
}

func NewHijriService(hijriRepository repository.HijriRepository) HijriService {
	return &hijriService{
		hijriRepository: hijriRepository,
	}
}
func (service *hijriService) GregorianToHijri(ctx context.Context, date, month, year int) (entity.Hijri, error) {
	var hijri entity.Hijri

	hijri, err := service.hijriRepository.GregorianToHijriCache(ctx, date, month, year)
	if err != nil {
		return hijri, err
	}

	if hijri.Date == "" {
		hijri, err = service.hijriRepository.GregorianToHijri(ctx, date, month, year)
		if err != nil {
			return hijri, err
		}

		if err = service.hijriRepository.SetGregorianToHijriCache(ctx, date, month, year, hijri); err != nil {
			return hijri, err
		}
	}

	return hijri, nil
}
