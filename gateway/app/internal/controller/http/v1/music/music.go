package music

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/metrics"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/music/dto"
	"github.com/VrMolodyakov/vgm/gateway/internal/controller/http/v1/validator"
	"github.com/VrMolodyakov/vgm/gateway/internal/domain/music/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"github.com/go-chi/chi"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	tracer = otel.Tracer("album-http")
)

type AlbumService interface {
	CreateAlbum(ctx context.Context, album model.Album) error
	FindAllAlbums(
		ctx context.Context,
		pagination model.Pagination,
		titleView model.AlbumTitleView,
		releaseView model.AlbumReleasedView,
		sort model.Sort) ([]model.AlbumPreview, error)
	CreatePerson(context.Context, model.Person) error
	FindFullAlbum(ctx context.Context, id string) (model.FullAlbum, error)
	FindLastUpdateDays(ctx context.Context, count uint64) ([]int64, error)
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
	ctx, span := tracer.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()

	var album dto.AlbumReq
	logger := logging.LoggerFromContext(ctx)
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	errs := validator.Validate(album)
	if errs != nil {
		jsonErr, _ := json.Marshal(errs)
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	}

	err := a.service.CreateAlbum(ctx, model.AlbumFromDto(album))
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
	ctx, span := tracer.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()

	var person dto.Person
	logger := logging.LoggerFromContext(ctx)
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	errs := validator.Validate(person)
	if errs != nil {
		jsonErr, _ := json.Marshal(errs)
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	}

	err := a.service.CreatePerson(ctx, model.PersonFromDto(person))
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
	ctx, span := tracer.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
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

	albums, err := a.service.FindAllAlbums(ctx, p, titleView, releaseView, s)
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
	response := make([]dto.AlbumPreviewRes, len(albums))
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
	ctx, span := tracer.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()

	albumID := chi.URLParam(r, "albumID")
	logger := logging.LoggerFromContext(ctx)
	fullAlbum, err := a.service.FindFullAlbum(ctx, albumID)
	if err != nil {
		logger.Error(err.Error())
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				metrics.AlbumCounter.WithLabelValues(albumID, strconv.Itoa(http.StatusInternalServerError))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			default:
				metrics.AlbumCounter.WithLabelValues(albumID, strconv.Itoa(http.StatusBadRequest))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	dto := fullAlbum.DtoFromModel()
	jsonResponse, err := json.Marshal(dto)
	if err != nil {
		metrics.AlbumCounter.WithLabelValues(albumID, strconv.Itoa(http.StatusInternalServerError))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	metrics.AlbumCounter.WithLabelValues(albumID, strconv.Itoa(http.StatusOK))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func (a *albumHandler) FindLastUpdateDays(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()

	param := chi.URLParam(r, "limit")
	logger := logging.LoggerFromContext(ctx)
	limit, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dates, err := a.service.FindLastUpdateDays(ctx, limit)
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
	jsonResponse, err := json.Marshal(dates)
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
	if field != "released_at" && field != "title" && field != "created_at" {
		return "", errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
	}

	return sortBy, nil
}
