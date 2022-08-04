package entity

type ReverseGeocode struct {
	Latitude                  float64 `json:"latitude"`
	Longitude                 float64 `json:"longitude"`
	Continent                 string  `json:"continent"`
	LookupSource              string  `json:"lookupSource"`
	ContinentCode             string  `json:"continentCode"`
	LocalityLanguageRequested string  `json:"localityLanguageRequested"`
	City                      string  `json:"city"`
	CountryName               string  `json:"countryName"`
	CountryCode               string  `json:"countryCode"`
	Postcode                  string  `json:"postcode"`
	PrincipalSubdivision      string  `json:"principalSubdivision"`
	PrincipalSubdivisionCode  string  `json:"principalSubdivisionCode"`
	PlusCode                  string  `json:"plusCode"`
	Locality                  string  `json:"locality"`
	LocalityInfo              struct {
		Administrative []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			IsoName     string `json:"isoName,omitempty"`
			Order       int    `json:"order"`
			AdminLevel  int    `json:"adminLevel"`
			IsoCode     string `json:"isoCode,omitempty"`
			WikidataID  string `json:"wikidataId"`
			GeonameID   int    `json:"geonameId"`
		} `json:"administrative"`
		Informative []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Order       int    `json:"order"`
			WikidataID  string `json:"wikidataId,omitempty"`
			GeonameID   int    `json:"geonameId,omitempty"`
			IsoName     string `json:"isoName,omitempty"`
			IsoCode     string `json:"isoCode,omitempty"`
		} `json:"informative"`
	} `json:"localityInfo"`
}
