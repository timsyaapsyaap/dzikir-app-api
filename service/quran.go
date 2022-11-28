package service

import (
	"context"
	"sync"

	"github.com/fahmialfareza/dzikir-app-api/entity"
	"github.com/fahmialfareza/dzikir-app-api/repository"
)

type QuranService interface {
	AllChapters(ctx context.Context) ([]entity.Chapter, error)
	AllVerses(ctx context.Context) (map[int][]entity.Verse, error)
	VersesByChapter(ctx context.Context, chapter int) ([]entity.Verse, error)
}

type quranService struct {
	quranRepository repository.QuranRepository
}

func NewQuranService(quranRepository repository.QuranRepository) QuranService {
	return &quranService{
		quranRepository: quranRepository,
	}
}

func (service *quranService) AllChapters(ctx context.Context) ([]entity.Chapter, error) {
	var chapters []entity.Chapter

	chapters, err := service.quranRepository.AllChaptersCache(ctx)
	if err != nil {
		return chapters, err
	}

	if len(chapters) == 0 {
		chapters, err = service.quranRepository.AllChapters(ctx)
		if err != nil {
			return chapters, err
		}

		if err = service.quranRepository.SetAllChaptersCache(ctx, chapters); err != nil {
			return chapters, err
		}
	}

	return chapters, nil
}

func (service *quranService) VersesByChapter(ctx context.Context, chapter int) ([]entity.Verse, error) {
	var verses []entity.Verse

	chapterDetails, err := service.quranRepository.GetChapterCache(ctx, chapter)
	if err != nil {
		return verses, err
	}

	if chapterDetails.ID == 0 {
		chapterDetails, err = service.quranRepository.GetChapter(ctx, chapter)
		if err != nil {
			return verses, err
		}

		if err = service.quranRepository.SetGetChapterCache(ctx, chapter, chapterDetails); err != nil {
			return verses, err
		}
	}

	// Get verse from redis
	verses, err = service.quranRepository.VersesByChapterCache(ctx, chapter)
	if err != nil {
		return verses, err
	}

	if len(verses) == 0 {
		verses, err = service.quranRepository.VersesByChapter(ctx, chapter, chapterDetails.VersesCount)
		if err != nil {
			return verses, err
		}

		if err = service.quranRepository.SetVersesByChapterCache(ctx, chapter, verses); err != nil {
			return verses, err
		}
	}

	return verses, nil
}

func (service *quranService) AllVerses(ctx context.Context) (map[int][]entity.Verse, error) {
	var (
		wg          sync.WaitGroup
		mutex       sync.Mutex
		totalVerses = 114
	)

	data := make(map[int][]entity.Verse)

	wg.Add(totalVerses)
	for i := 1; i <= totalVerses; i++ {
		go func(i int) error {
			defer mutex.Unlock()
			defer wg.Done()

			chapterDetails, err := service.quranRepository.GetChapter(ctx, i)
			if err != nil {
				return err
			}

			versesDetails, err := service.quranRepository.VersesByChapter(ctx, i, chapterDetails.VersesCount)
			if err != nil {
				return err
			}

			mutex.Lock()
			data[i] = versesDetails

			return nil
		}(i)
	}
	wg.Wait()

	return data, nil
}
