package repository

import "fmt"

const (
	findCityAddress = "/v1/sholat/kota/cari/"
	cityDetails     = "/v1/sholat/kota/id/"
	allCities       = "/v1/sholat/kota/semua"
)

func schedule(cityId int, year int, month int, date int) string {
	return fmt.Sprintf("/v1/sholat/jadwal/%v/%v/%v/%v", cityId, year, month, date)
}
