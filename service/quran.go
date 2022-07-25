package service

import (
	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
)

type QuranService interface {
	AllChapters() ([]entity.Chapter, error)
	VersesByChapter(chapter int) ([]entity.Verse, error)
}

type quranService struct {
	quranRepository repository.QuranRepository
}

func NewQuranService(quranRepository repository.QuranRepository) QuranService {
	return &quranService{
		quranRepository: quranRepository,
	}
}

func (service *quranService) AllChapters() ([]entity.Chapter, error) {
	return service.quranRepository.AllChapters()
}

func (service *quranService) VersesByChapter(chapter int) ([]entity.Verse, error) {
	chapterDetails, err := service.quranRepository.GetChapter(chapter)
	if err != nil {
		return nil, err
	}

	return service.quranRepository.VersesByChapter(chapter, chapterDetails.VersesCount)
}
