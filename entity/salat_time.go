package entity

// These stuct is getting data from API
// Salat Time represents a time of salat.
type SalatTimeRestAPI struct {
	ID         string                     `json:"id"`
	City       string                     `json:"lokasi"`
	Province   string                     `json:"daerah"`
	Coordinate SalatTimeCoordinateRestAPI `json:"koordinat"`
	Schedule   SalatTimeScheduleRestAPI   `json:"jadwal"`
}

// Coordinate defines a coordinate of salat time.
type SalatTimeCoordinateRestAPI struct {
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lon"`
	Latitude  string  `json:"lintang"`
	Longitude string  `json:"bujur"`
}

type SalatTimeScheduleRestAPI struct {
	FullDate string `json:"tanggal"`
	Imsak    string `json:"imsak"`
	Fajr     string `json:"subuh"`
	Rise     string `json:"terbit"`
	Duha     string `json:"dhuha"`
	Dhuhr    string `json:"dzuhur"`
	Asr      string `json:"ashar"`
	Maghrib  string `json:"maghrib"`
	Isha     string `json:"isya"`
	Date     string `json:"date"`
}

// SalatTimeCity represents a city of salat time.
type SalatTimeCityRestAPI struct {
	ID   string `json:"id"`
	City string `json:"lokasi"`
}

type SalatTimeCityFindRestAPIResponse struct {
	Status bool                   `json:"status"`
	Data   []SalatTimeCityRestAPI `json:"data"`
}

type SalatTimeCityDetailsRestAPIResponse struct {
	Status bool                 `json:"status"`
	Data   SalatTimeCityRestAPI `json:"data"`
}

type SalatTimeRestAPIResponse struct {
	Status bool             `json:"status"`
	Data   SalatTimeRestAPI `json:"data"`
}

// SalatTime that returned to user
// Salat Time represents a time of salat.
type SalatTime struct {
	ID         int                 `json:"id"`
	City       string              `json:"city"`
	Province   string              `json:"province"`
	Coordinate SalatTimeCoordinate `json:"coordinate"`
	Schedule   SalatTimeSchedule   `json:"schedule"`
}

// Coordinate defines a coordinate of salat time.
type SalatTimeCoordinate struct {
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Latitude  string  `json:"latitude"`
	Longitude string  `json:"longitude"`
}

type SalatTimeSchedule struct {
	FullDate string `json:"full_date"`
	Imsak    string `json:"imsak"`
	Fajr     string `json:"fajr"`
	Rise     string `json:"rise"`
	Duha     string `json:"duha"`
	Dhuhr    string `json:"dhuhr"`
	Asr      string `json:"asr"`
	Maghrib  string `json:"maghrib"`
	Isha     string `json:"isha"`
	Date     string `json:"date"`
}

// SalatTimeCity represents a city of salat time.
type SalatTimeCity struct {
	ID   int    `json:"id"`
	City string `json:"city"`
}
