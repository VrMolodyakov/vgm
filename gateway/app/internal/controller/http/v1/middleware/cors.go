package middleware

import "net/http"

type Cors struct {
	AllowOrigins   string
	AllowHeaders   string
	AllowedMethods string
}

func NewCors(allowOrigins string, allowHeaders string, allowedMethods string) Cors {
	return Cors{
		AllowOrigins:   allowOrigins,
		AllowHeaders:   allowHeaders,
		AllowedMethods: allowedMethods,
	}
}

func (c *Cors) CORS(next http.Handler) http.Handler {
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
