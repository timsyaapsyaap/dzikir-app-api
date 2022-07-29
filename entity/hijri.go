package entity

type HijriReponseAPI struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		Hijri     Hijri `json:"hijri"`
		Gregorian struct {
			Date    string `json:"date"`
			Format  string `json:"format"`
			Day     string `json:"day"`
			Weekday struct {
				En string `json:"en"`
			} `json:"weekday"`
			Month struct {
				Number int    `json:"number"`
				En     string `json:"en"`
			} `json:"month"`
			Year        string `json:"year"`
			Designation struct {
				Abbreviated string `json:"abbreviated"`
				Expanded    string `json:"expanded"`
			} `json:"designation"`
		} `json:"gregorian"`
	} `json:"data"`
}

type Hijri struct {
	Date    string `json:"date"`
	Format  string `json:"format"`
	Day     string `json:"day"`
	Weekday struct {
		En string `json:"en"`
		Ar string `json:"ar"`
	} `json:"weekday"`
	Month struct {
		Number int    `json:"number"`
		En     string `json:"en"`
		Ar     string `json:"ar"`
		Id     string `json:"id"`
	} `json:"month"`
	Year        string `json:"year"`
	Designation struct {
		Abbreviated string `json:"abbreviated"`
		Expanded    string `json:"expanded"`
	} `json:"designation"`
	Holidays []interface{} `json:"holidays"`
}
