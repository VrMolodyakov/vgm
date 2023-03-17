package middleware

import "net/http"

type cors struct {
	AllowOrigins   string
	AllowHeaders   string
	AllowedMethods string
}

func NewCors(allowOrigins string, allowHeaders string, allowedMethods string) *cors {
	return &cors{
		AllowOrigins:   allowOrigins,
		AllowHeaders:   allowHeaders,
		AllowedMethods: allowedMethods,
	}
}

func (c *cors) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", c.AllowOrigins)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", c.AllowHeaders)
		w.Header().Set("Access-Control-Allow-Methods", c.AllowedMethods)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})

}
