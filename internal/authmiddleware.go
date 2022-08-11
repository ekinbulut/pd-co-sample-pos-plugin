package internal

import (
	"net/http"
	"pos-plugin/internal/response"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get auth token from header
		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write(response.CreateAErrorResponse("missing auth token"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
