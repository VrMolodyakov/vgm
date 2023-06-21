package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/VrMolodyakov/vgm/youtube/pkg/logging"
	"github.com/go-chi/chi"
	"go.opentelemetry.io/otel"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	tracer          = otel.Tracer("youtube-http")
	playlist string = "https://www.youtube.com/watch_videos?video_ids="
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
	limitStr := chi.URLParam(r, "limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	titles, err := y.music.FindRandomTitles(ctx, uint64(limit))
	if err != nil {
		y.logger.Error(err.Error())
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
	g, ctx := errgroup.WithContext(ctx)
	ids := make([]string, len(titles))
	for i, title := range titles {
		i, title := i, title
		g.Go(func() error {
			id, err := y.youtube.GetVideoIDByTitle(ctx, title)
			if err == nil {
				ids[i] = id
			}
			return err
		})

	}
	if err := g.Wait(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if empty := emptyPlaylist(ids); empty {
		http.Error(w, "cannot collect video IDs", http.StatusNotFound)
		return
	}
	resURL := createPlaylistUrl(playlist, ids)
	var url CreatePlaylistRes
	url.URL = resURL
	jsonResponse, err := json.Marshal(url)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func createPlaylistUrl(playlistBase string, ids []string) string {
	var sb strings.Builder
	n := len(ids) + len(playlistBase)
	for i := range ids {
		n += len(ids[i])
	}
	sb.Grow(n)
	sb.WriteString(playlistBase)
	var i int
	for ; i < len(ids); i++ {
		if ids[i] != "" {
			sb.WriteString(ids[i])
			break
		}
	}
	for i := i + 1; i < len(ids); i++ {
		if ids[i] != "" {
			sb.WriteString(",")
			sb.WriteString(ids[i])
		}
	}
	return sb.String()
}

func emptyPlaylist(ids []string) bool {
	for _, id := range ids {
		if id != "" {
			return false
		}
	}
	return true
}
