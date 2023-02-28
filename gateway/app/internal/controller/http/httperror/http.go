package httperror

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Ok      bool   `json:"ok"`
	Message string `json:"err"`
}

func NewError(w http.ResponseWriter, code int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(
		Error{
			Ok:      false,
			Message: err.Error(),
		},
	)
}
