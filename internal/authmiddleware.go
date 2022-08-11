package internal

import (
	"log"
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

		// log auth token
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		log.Printf("auth token: %s", authToken)

		next.ServeHTTP(w, r)
	})
}
