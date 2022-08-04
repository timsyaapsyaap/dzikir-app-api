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

func monthIdHijri(month int) string {
	switch month {
	case 1:
		return "Muharram"
	case 2:
		return "Safar"
	case 3:
		return "Rabiul-Awwal"
	case 4:
		return "Rabiul-Akhir"
	case 5:
		return "Jumadil-Awwal"
	case 6:
		return "Jumadil-Akhir"
	case 7:
		return "Rajab"
	case 8:
		return "Sya'ban"
	case 9:
		return "Ramadhan"
	case 10:
		return "Shawwal"
	case 11:
		return "Dhul-Qa'dah"
	case 12:
		return "Dhul-Hijjah"
	}

	return ""
}

func gregorianToHijri(date int, month int, year int) string {
	return fmt.Sprintf("/v1/gToH?date=%v-%v-%v", date, month, year)
}

func reverseGeocode(lat float64, lng float64) string {
	return fmt.Sprintf(`/data/reverse-geocode-client?latitude=%v&longitude=%v`, lat, lng)
}
