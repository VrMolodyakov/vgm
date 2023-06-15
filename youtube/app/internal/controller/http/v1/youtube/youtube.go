package youtube

import "net/http"

type youtubeHandler struct {
}

func New() *youtubeHandler {
	return &youtubeHandler{}
}

func (y *youtubeHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
}
