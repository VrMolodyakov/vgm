package album

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

	err := a.service.CreateAlbum(r.Context(), model.AlbumFromDto(album))
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

func (a *albumHandler) FindAllAlbums(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sort_by")
	sortBy, err := parseSortQuery(sortBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	strLimit := r.URL.Query().Get("limit")
	limit := -1
	if strLimit != "" {
		limit, err = strconv.Atoi(strLimit)
		if err != nil || limit <= -1 {
			http.Error(w, "limit query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}
	strOffset := r.URL.Query().Get("offset")
	offset := -1
	if strOffset != "" {
		offset, err = strconv.Atoi(strOffset)
		if err != nil || offset <= -1 {
			http.Error(w, "offset query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}
	fildFilterVal := r.URL.Query().Get("title.val")
	fildFilterOp := r.URL.Query().Get("title.op")

	fmt.Println(fildFilterOp)
	fmt.Println(fildFilterVal)

	// client sirvice call here
	w.Header().Add("Content-Type", "application/json")
	// if err := json.NewEncoder(w).Encode(users); err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	// }
}

//title
//released_at

// Filter for string values, example: ?email.op=eq&email.val=me@example.com
/*

{
    "pagination": {
        "limit": "4",
        "offset": "20539"
    },
    "released_at": {
        "op": 2,
        "val": "deserunt magna in amet"
    },
    "sort": {
        "field": "elit laborum in aliqua"
    },
    "title": {
        "op": "OPERATOR_LIKE",
        "val": "eu aliquip"
    }
}


*/

func parseSortQuery(sortBy string) (string, error) {
	if sortBy == "" {
		return sortBy, nil
	}
	field := strings.TrimPrefix(sortBy, "-")
	if field != "released_at" && field != "title" {
		return "", errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
	}

	return sortBy, nil
}
