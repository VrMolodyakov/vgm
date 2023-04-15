package dto

type Album struct {
	Album     AlbumView `json:"album"`
	Info      Info      `json:"info"`
	Tracklist []Track   `json:"tracklist"`
	Credits   []Credit  `json:"credits"`
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate int64  `json:"birth_date"`
}

type AlbumView struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	ReleasedAt int64  `json:"released_at"`
	CreatedAt  int64  `json:"created_at"`
}

type Info struct {
	ID             string  `json:"id"`
	AlbumID        string  `json:"album_id"`
	CatalogNumber  string  `json:"catalog_number"`
	FullImageSrc   string  `json:"full_image_src"`
	SmallImageSrc  string  `json:"small_image_src"`
	Barcode        string  `json:"barcode"`
	CurrencyCode   string  `json:"currency_code"`
	MediaFormat    string  `json:"media_format"`
	Classification string  `json:"classification"`
	Publisher      string  `json:"publisher"`
	Price          float64 `json:"price"`
}

type Track struct {
	ID       int64  `json:"id"`
	AlbumID  string `json:"album_id"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
}

type Credit struct {
	PersonID   int64  `json:"person_id"`
	AlbumID    string `json:"album_id"`
	Profession string `json:"profession"`
}
