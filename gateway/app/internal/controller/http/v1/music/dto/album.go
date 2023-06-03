package dto

type AlbumReq struct {
	Album     AlbumViewReq `json:"album"`
	Info      InfoReq      `json:"info"`
	Tracklist []TrackReq   `json:"tracklist"`
	Credits   []CreditReq  `json:"credits"`
}

type Person struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	BirthDate int64  `json:"birth_date" validate:"required"`
}

type AlbumViewReq struct {
	Title      string `json:"title" validate:"required"`
	ReleasedAt int64  `json:"released_at" validate:"gt=0,required,numeric"`
	CreatedAt  int64  `json:"created_at"`
}

type AlbumViewRes struct {
	AlbumID    string `json:"album_id"`
	Title      string `json:"title"`
	ReleasedAt int64  `json:"released_at"`
	CreatedAt  int64  `json:"created_at"`
}

type AlbumPreviewRes struct {
	AlbumID       string `json:"album_id"`
	Title         string `json:"title"`
	ReleasedAt    int64  `json:"released_at"`
	CreatedAt     int64  `json:"created_at"`
	Publisher     string `json:"publisher" validate:"required"`
	SmallImageSrc string `json:"small_image_src"`
}

type InfoReq struct {
	CatalogNumber  string  `json:"catalog_number" validate:"required"`
	FullImageSrc   string  `json:"full_image_src"`
	SmallImageSrc  string  `json:"small_image_src"`
	Barcode        string  `json:"barcode" validate:"required"`
	CurrencyCode   string  `json:"currency_code" validate:"required"`
	MediaFormat    string  `json:"media_format" validate:"required"`
	Classification string  `json:"classification" validate:"required"`
	Publisher      string  `json:"publisher" validate:"required"`
	Price          float64 `json:"price" validate:"gt=-1"`
}

type InfoRes struct {
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
	Title    string `json:"title" validate:"required,min=1"`
	Duration string `json:"duration" validate:"required,min=3"`
}

type TrackRes struct {
	ID       int64  `json:"id"`
	AlbumID  string `json:"album_id"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
}

type CreditReq struct {
	PersonID   int64  `json:"person_id" validate:"gt=0,required"`
	Profession string `json:"profession" validate:"required,min=3"`
}

type CreditInfoRes struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Profession string `json:"profession"`
}

type FullAlbumReq struct {
	Id string `json:"album_id"  validate:"required"`
}

type FullAlbumRes struct {
	Album     AlbumViewRes    `json:"album"`
	Info      InfoRes         `json:"info"`
	Tracklist []TrackRes      `json:"tracklist"`
	Credits   []CreditInfoRes `json:"credits"`
}

type DatesRes struct {
	Dates []int64 `json:"dates"`
}
