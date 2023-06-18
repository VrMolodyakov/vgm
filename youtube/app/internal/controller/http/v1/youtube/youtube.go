package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VrMolodyakov/vgm/youtube/internal/controller/http/v1/validator"
	"github.com/VrMolodyakov/vgm/youtube/pkg/logging"
	"go.opentelemetry.io/otel"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	fmt.Println("here")
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

	titles, err := y.music.FindRandomTitles(ctx, req.Count)
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
	playlistTitle := fmt.Sprintf("random VGM playlist date = %s", generateDateString())
	playlistID, err := y.youtube.CreatePlaylist(playlistTitle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	y.youtube.AddVideosToPlaylist(playlistID, ids)
	jsonResponse, err := json.Marshal(titles)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func generateDateString() string {
	currentTime := time.Now()
	dateString := currentTime.Format("2006-01-02 15:04:05")
	return dateString
}
