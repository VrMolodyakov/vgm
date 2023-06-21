package youtube

type CreatePlaylistReq struct {
	Count uint64 `json:"count"`
}

type CreatePlaylistRes struct {
	URL string `json:"url"`
}
