package dto

type AlbumReq struct {
	Album     AlbumViewReq `json:"album"`
	Info      InfoReq      `json:"info"`
	Tracklist []TrackReq   `json:"tracklist"`
	Credits   []CreditReq  `json:"credits"`
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate int64  `json:"birth_date"`
}

type AlbumViewReq struct {
	Title      string `json:"title"`
	ReleasedAt int64  `json:"released_at"`
	CreatedAt  int64  `json:"created_at"`
}

type AlbumViewRes struct {
	AlbumID    string `json:"album_id"`
	Title      string `json:"title"`
	ReleasedAt int64  `json:"released_at"`
	CreatedAt  int64  `json:"created_at"`
}

type InfoReq struct {
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

type InfoRes struct {
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

type TrackReq struct {
	Title    string `json:"title"`
	Duration string `json:"duration"`
}

type TrackRes struct {
	ID       int64  `json:"id"`
	AlbumID  string `json:"album_id"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
}

type CreditReq struct {
	PersonID   int64  `json:"person_id"`
	Profession string `json:"profession"`
}

type CreditInfoRes struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Profession string `json:"profession"`
}

type FullAlbumReq struct {
	Id string `json:"album_id"`
}

type FullAlbumRes struct {
	Album     AlbumViewRes    `json:"album"`
	Info      InfoRes         `json:"info"`
	Tracklist []TrackRes      `json:"tracklist"`
	Credits   []CreditInfoRes `json:"credits"`
}
