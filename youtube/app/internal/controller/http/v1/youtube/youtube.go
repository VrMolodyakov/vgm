package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VrMolodyakov/vgm/youtube/internal/controller/http/v1/validator"
	"github.com/VrMolodyakov/vgm/youtube/pkg/logging"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("youtube-http")
)

type youtubeHandler struct {
	logger  logging.Logger
	youtube YoutubeService
	music   MusicService
}

func NewYoutubeHandler(
	logger logging.Logger,
	youtube YoutubeService,
	music MusicService,

) *youtubeHandler {
	return &youtubeHandler{
		music:   music,
		youtube: youtube,
		logger:  logger,
	}
}

func (y *youtubeHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))
	defer span.End()

	var req CreatePlaylistReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	errs := validator.Validate(req)
	if errs != nil {
		jsonErr, _ := json.Marshal(errs)
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	}

	y.music.FindRandomTitles(ctx, req.Count)
}
