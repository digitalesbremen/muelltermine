package api

import (
	"encoding/json"
	"net/http"
)

func Handle404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

		_ = json.
			NewEncoder(w).
			Encode(
				protocolError{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				})
	})
}

type protocolError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
