package service

import (
	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
)

type HijriService interface {
	GregorianToHijri(date int, month int, year int) (entity.Hijri, error)
}

type hijriService struct {
	hijriRepository repository.HijriRepository
}

func NewHijriService(hijriRepository repository.HijriRepository) HijriService {
	return &hijriService{
		hijriRepository: hijriRepository,
	}
}
func (service *hijriService) GregorianToHijri(date int, month int, year int) (entity.Hijri, error) {
	return service.hijriRepository.GregorianToHijri(date, month, year)
}
