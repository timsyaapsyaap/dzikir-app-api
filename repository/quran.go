package repository

import (
	"encoding/json"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/helper"
)

type QuranRepository interface {
	AllChapters() ([]entity.Chapter, error)
	VersesByChapter(chapter int, perPage int) ([]entity.Verse, error)
	GetChapter(chapter int) (entity.Chapter, error)
}

type quranRepository struct {
	api *entity.Config
}

func NewQuranRepository(api *entity.Config) QuranRepository {
	return &quranRepository{
		api: api,
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
