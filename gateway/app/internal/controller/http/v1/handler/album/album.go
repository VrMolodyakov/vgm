package album

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/handler/album/dto"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/music/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"github.com/go-chi/chi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AlbumService interface {
	CreateAlbum(ctx context.Context, album model.Album) error
	FindAllAlbums(
		ctx context.Context,
		pagination model.Pagination,
		titleView model.AlbumTitleView,
		releaseView model.AlbumReleasedView,
		sort model.Sort) ([]model.AlbumView, error)
	CreatePerson(context.Context, model.Person) error
	FindFullAlbum(ctx context.Context, id string) (model.FullAlbum, error)
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
	logger := logging.LoggerFromContext(r.Context())
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}
	err := a.service.CreateAlbum(r.Context(), model.AlbumFromDto(album))
	if err != nil {
		logger.Error(err.Error())
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			default:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("the album was created"))
}

//TODO:return person?
func (a *albumHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person dto.Person
	logger := logging.LoggerFromContext(r.Context())
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}
	err := a.service.CreatePerson(r.Context(), model.PersonFromDto(person))
	if err != nil {
		logger.Error(err.Error())
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			default:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("the person was created"))
}

func (a *albumHandler) FindAllAlbums(w http.ResponseWriter, r *http.Request) {
	logger := logging.LoggerFromContext(r.Context())
	sortBy := r.URL.Query().Get("sort_by")
	sortBy, err := validateSortQuery(sortBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	strLimit := r.URL.Query().Get("limit")
	limit := 0
	if strLimit != "" {
		limit, err = strconv.Atoi(strLimit)
		if err != nil || limit <= -1 {
			http.Error(w, "limit query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}
	strOffset := r.URL.Query().Get("offset")
	offset := 0
	if strOffset != "" {
		offset, err = strconv.Atoi(strOffset)
		if err != nil || offset <= -1 {
			http.Error(w, "offset query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}
	titlFilterVal := r.URL.Query().Get("title.val")
	titleFilterOp := r.URL.Query().Get("title.op")
	releaseFilterVal := r.URL.Query().Get("release.val")
	releaseFilterOp := r.URL.Query().Get("release.op")

	titleView := model.AlbumTitleView{
		Value:    titlFilterVal,
		Operator: titleFilterOp,
	}

	releaseView := model.AlbumReleasedView{
		Value:    releaseFilterVal,
		Operator: releaseFilterOp,
	}

	p := model.Pagination{
		Limit:  limit,
		Offset: offset,
	}

	s := model.Sort{
		Field: sortBy,
	}

	albums, err := a.service.FindAllAlbums(r.Context(), p, titleView, releaseView, s)
	if err != nil {
		logger.Error(err.Error())
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			default:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	response := make([]dto.AlbumView, len(albums))
	for i := 0; i < len(response); i++ {
		response[i] = albums[i].DtoFromModel()
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (a *albumHandler) FindFullAlbums(w http.ResponseWriter, r *http.Request) {
	// var req dto.FullAlbumRequest
	albumID := chi.URLParam(r, "albumID")
	logger := logging.LoggerFromContext(r.Context())
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
	// 	return
	// }
	fullAlbum, err := a.service.FindFullAlbum(r.Context(), albumID)
	if err != nil {
		logger.Error(err.Error())
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			default:
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	jsonResponse, err := json.Marshal(fullAlbum)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func validateSortQuery(sortBy string) (string, error) {
	if sortBy == "" {
		return sortBy, nil
	}
	field := strings.TrimPrefix(sortBy, "-")
	if field != "released_at" && field != "title" {
		return "", errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
	}

	return sortBy, nil
}
