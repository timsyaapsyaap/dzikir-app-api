package repository

import (
	"encoding/json"
	"fmt"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/gomodule/redigo/redis"
)

type QuranRepository interface {
	AllChapters() ([]entity.Chapter, error)
	VersesByChapter(chapter int, perPage int) ([]entity.Verse, error)
	VersesByChapterRedis(chapter int) ([]entity.Verse, error)
	GetChapter(chapter int) (entity.Chapter, error)
}

type quranRepository struct {
	api   *entity.Config
	redis *redis.Pool
}

func NewQuranRepository(api *entity.Config, redis *redis.Pool) QuranRepository {
	return &quranRepository{
		api:   api,
		redis: redis,
	}
}

func (repository *quranRepository) AllChapters() ([]entity.Chapter, error) {
	var (
		dataResponse entity.ChaptersRestAPI
		data         []entity.Chapter
	)

	body, err := helper.GetRequest(repository.api.QuranRestApi + allChapters)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	return dataResponse.Chapters, nil
}

func (repository *quranRepository) VersesByChapter(chapter int, perPage int) ([]entity.Verse, error) {
	var (
		dataResponse entity.VersesRestAPI
		data         []entity.Verse
	)

	body, err := helper.GetRequest(repository.api.QuranRestApi + versesByChapter(chapter, perPage))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	return dataResponse.Verses, nil
}

func (repository *quranRepository) VersesByChapterRedis(chapter int) ([]entity.Verse, error) {
	var (
		data     []entity.Verse
		verseKey = "verse:" + fmt.Sprint(chapter)
	)

	// Get verse from redis
	client := repository.redis.Get()
	defer client.Close()

	verse, err := client.Do("GET", verseKey)
	if err != nil {
		return nil, err
	} else if verse == nil {
		return nil, err
	}

	err = json.Unmarshal(verse.([]byte), &data)

	return data, err
}

func (repository *quranRepository) GetChapter(chapter int) (entity.Chapter, error) {
	var (
		dataResponse entity.ChapterRestAPI
		data         entity.Chapter
	)

	body, err := helper.GetRequest(repository.api.QuranRestApi + getChapter(chapter))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	return dataResponse.Chapter, nil
}
