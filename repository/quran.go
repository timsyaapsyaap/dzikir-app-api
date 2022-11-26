package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/helper"
	"github.com/go-redis/redis/v8"
)

type QuranRepository interface {
	AllChapters(ctx context.Context) ([]entity.Chapter, error)
	VersesByChapter(ctx context.Context, chapter, perPage int) ([]entity.Verse, error)
	GetChapter(ctx context.Context, chapter int) (entity.Chapter, error)

	AllChaptersCache(ctx context.Context) ([]entity.Chapter, error)
	SetAllChaptersCache(ctx context.Context, chapters []entity.Chapter) error
	GetChapterCache(ctx context.Context, chapter int) (entity.Chapter, error)
	SetGetChapterCache(ctx context.Context, chapter int, chapterData entity.Chapter) error
	VersesByChapterCache(ctx context.Context, chapter int) ([]entity.Verse, error)
	SetVersesByChapterCache(ctx context.Context, chapter int, verses []entity.Verse) error
}

type quranRepository struct {
	api         *entity.Config
	redisClient *redis.Client
}

func NewQuranRepository(api *entity.Config, redisClient *redis.Client) QuranRepository {
	return &quranRepository{
		api:         api,
		redisClient: redisClient,
	}
}

func (repository *quranRepository) AllChapters(ctx context.Context) ([]entity.Chapter, error) {
	var (
		dataResponse entity.ChaptersRestAPI
		data         []entity.Chapter
	)

	body, err := helper.GetRequest(ctx, repository.api.QuranRestApi+allChapters)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	return dataResponse.Chapters, nil
}

func (repository *quranRepository) VersesByChapter(ctx context.Context, chapter, perPage int) ([]entity.Verse, error) {
	var (
		dataResponse entity.VersesRestAPI
		data         []entity.Verse
	)

	body, err := helper.GetRequest(ctx, repository.api.QuranRestApi+versesByChapter(chapter, perPage))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	return dataResponse.Verses, nil
}

func (repository *quranRepository) GetChapter(ctx context.Context, chapter int) (entity.Chapter, error) {
	var (
		dataResponse entity.ChapterRestAPI
		data         entity.Chapter
	)

	body, err := helper.GetRequest(ctx, repository.api.QuranRestApi+getChapter(chapter))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return data, err
	}

	return dataResponse.Chapter, nil
}

func (repository *quranRepository) AllChaptersCache(ctx context.Context) ([]entity.Chapter, error) {
	var data []entity.Chapter

	// Get verse from redis
	chapters, err := repository.redisClient.Get(ctx, allChaptersKey).Result()
	if err == redis.Nil {
		return data, nil
	} else if err != nil {
		return data, err
	}

	if err = json.Unmarshal([]byte(chapters), &data); err != nil {
		return data, err
	}

	return data, err
}

func (repository *quranRepository) SetAllChaptersCache(ctx context.Context, chapters []entity.Chapter) error {
	chaptersByte, err := json.Marshal(chapters)
	if err != nil {
		return err
	}

	if err = repository.redisClient.Set(ctx, allChaptersKey, string(chaptersByte), redisQuranExpiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repository *quranRepository) VersesByChapterCache(ctx context.Context, chapter int) ([]entity.Verse, error) {
	var (
		data []entity.Verse
		key  = fmt.Sprintf(verseKey, chapter)
	)

	// Get verse from redis
	verses, err := repository.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return data, nil
	} else if err != nil {
		return data, err
	}

	if err = json.Unmarshal([]byte(verses), &data); err != nil {
		return data, err
	}

	return data, err
}

func (repository *quranRepository) SetVersesByChapterCache(ctx context.Context, chapter int, verses []entity.Verse) error {
	key := fmt.Sprintf(verseKey, chapter)

	versesByte, err := json.Marshal(verses)
	if err != nil {
		return err
	}

	if err = repository.redisClient.Set(ctx, key, string(versesByte), redisQuranExpiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repository *quranRepository) GetChapterCache(ctx context.Context, chapter int) (entity.Chapter, error) {
	var (
		data entity.Chapter
		key  = fmt.Sprintf(chapterKey, chapter)
	)

	// Get verse from redis
	chapterData, err := repository.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return data, nil
	} else if err != nil {
		return data, err
	}

	if err = json.Unmarshal([]byte(chapterData), &data); err != nil {
		return data, err
	}

	return data, err
}

func (repository *quranRepository) SetGetChapterCache(ctx context.Context, chapter int, chapterData entity.Chapter) error {
	key := fmt.Sprintf(chapterKey, chapter)

	chapterDataByte, err := json.Marshal(chapterData)
	if err != nil {
		return err
	}

	if err = repository.redisClient.Set(ctx, key, string(chapterDataByte), redisQuranExpiration).Err(); err != nil {
		return err
	}

	return nil
}
