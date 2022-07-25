package entity

// These stuct is getting data from API
// Chapters lists from Rest API
type ChaptersRestAPI struct {
	Chapters []Chapter `json:"chapters"`
}

type ChapterRestAPI struct {
	Chapter Chapter `json:"chapter"`
}

// Verses lists from Rest API
type VersesRestAPI struct {
	Verses []Verse `json:"verses"`
}

// Quran that returned to user
// Chapter represents a chapter of quran.
type Chapter struct {
	ID              int            `json:"id"`
	RevelationPlace string         `json:"revelation_place"`
	RevelationOrder int            `json:"revelation_order"`
	BismillahPre    bool           `json:"bismillah_pre"`
	NameSimple      string         `json:"name_simple"`
	NameComplex     string         `json:"name_complex"`
	NameArabic      string         `json:"name_arabic"`
	VersesCount     int            `json:"verses_count"`
	Pages           []int          `json:"pages"`
	TranslatedName  TranslatedName `json:"translated_name"`
}

type TranslatedName struct {
	LanguageName string `json:"language_name"`
	Name         string `json:"name"`
}

type Verse struct {
	ID              int           `json:"id"`
	VerseNumber     int           `json:"verse_number"`
	VerseKey        string        `json:"verse_key"`
	JuzNumber       int           `json:"juz_number"`
	HizbNumber      int           `json:"hizb_number"`
	RubElHizbNumber int           `json:"rub_el_hizb_number"`
	RukuNumber      int           `json:"ruku_number"`
	ManzilNumber    int           `json:"manzil_number"`
	SajdahNumber    int           `json:"sajdah_number"`
	TextUthmani     string        `json:"text_uthmani"`
	TextIndopak     string        `json:"text_indopak"`
	PageNumber      int           `json:"page_number"`
	Audio           Audio         `json:"audio"`
	Translations    []Translation `json:"translations"`
}

type Audio struct {
	URL      string  `json:"url"`
	Segments [][]int `json:"segments"`
}

type Translation struct {
	ID         int    `json:"id"`
	ResourceID int    `json:"resource_id"`
	Text       string `json:"text"`
}
