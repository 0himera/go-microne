package httpapi

import (
	"net/http"

	"github.com/0himera/go-microne/internal/domains"
)


func RequestIDMiddleware(genID domains.IDGenerator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := genID.NewID()
			w.Header().Set(RequestIDHeader, requestID)
			next.ServeHTTP(w, r)
		})
	}
}