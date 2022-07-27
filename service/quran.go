package service

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
	"github.com/gomodule/redigo/redis"
)

type QuranService interface {
	AllChapters() ([]entity.Chapter, error)
	AllVerses() (map[int][]entity.Verse, error)
	VersesByChapter(chapter int) ([]entity.Verse, error)
}

type quranService struct {
	quranRepository repository.QuranRepository
	redis           *redis.Pool
}

func NewQuranService(quranRepository repository.QuranRepository, redis *redis.Pool) QuranService {
	return &quranService{
		quranRepository: quranRepository,
		redis:           redis,
	}
}

func (service *quranService) AllChapters() ([]entity.Chapter, error) {
	return service.quranRepository.AllChapters()
}

func (service *quranService) VersesByChapter(chapter int) ([]entity.Verse, error) {
	var (
		data     []entity.Verse
		verseKey = "verse:" + fmt.Sprint(chapter)
	)

	chapterDetails, err := service.quranRepository.GetChapter(chapter)
	if err != nil {
		return nil, err
	}

	// Get verse from redis
	verse, err := service.quranRepository.VersesByChapterRedis(chapter)

	if err != nil {
		return data, err
	} else if verse == nil {
		data, err = service.quranRepository.VersesByChapter(chapter, chapterDetails.VersesCount)
		if err != nil {
			return nil, err
		}

		// Save verse to redis
		client := service.redis.Get()
		defer client.Close()

		jsonData, err := json.Marshal(data)
		if err != nil {
			return data, err
		}

		// Save to redis
		_, err = client.Do("SET", verseKey, jsonData)

		return data, err
	}

	data = verse

	return data, nil
}

func (service *quranService) AllVerses() (map[int][]entity.Verse, error) {
	var wg sync.WaitGroup

	data := make(map[int][]entity.Verse)

	wg.Add(114)
	for i := 1; i <= 114; i++ {
		go func(i int) {
			defer wg.Done()

			chapterDetails, err := service.quranRepository.GetChapter(i)
			if err != nil {
				fmt.Println(err)
			}

			versesDetails, err := service.quranRepository.VersesByChapter(i, chapterDetails.VersesCount)
			if err != nil {
				fmt.Println(err)
			}

			data[i] = versesDetails
		}(i)
	}
	wg.Wait()

	return data, nil
}
