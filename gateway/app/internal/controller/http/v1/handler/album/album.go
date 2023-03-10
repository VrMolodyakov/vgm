package album

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/handler/album/dto"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
)

type AlbumService interface {
	CreateAlbum(ctx context.Context, album model.Album) error
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
	logger := logging.GetLogger()
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err := a.service.CreateAlbum(context.Background(), model.AlbumFromDto(album))
	if err != nil {
		logger.Error(err.Error())
		if _, ok := errors.IsInternal(err); ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("the album was created"))
}
