package repository

import "fmt"

const (
	findCityAddress = "/v1/sholat/kota/cari/"
	cityDetails     = "/v1/sholat/kota/id/"
	allCities       = "/v1/sholat/kota/semua"

	allChapters = "/api/v4/chapters?language=id"
)

func schedule(cityId int, year int, month int, date int) string {
	return fmt.Sprintf("/v1/sholat/jadwal/%v/%v/%v/%v", cityId, year, month, date)
}

func getChapter(id int) string {
	return fmt.Sprintf("/api/v4/chapters/%v?language=id", id)
}

func versesByChapter(chapter int, perPage int) string {
	return fmt.Sprintf("/api/v4/verses/by_chapter/%v?language=id&fields=text_uthmani,text_indopak&translations=33&audio=7&word_fields=id&per_page=%v", chapter, perPage)
}
