package album

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/handler/album/dto"
)

type AlbumService interface {
	CreateAlbum(context.Context)
}

type albumHandler struct {
	service AlbumService
}

func NewAlbumHandler(service AlbumService) *albumHandler {
	return &albumHandler{
		service: service,
	}
}

func (a *albumHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album dto.Album

	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

}
