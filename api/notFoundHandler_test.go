package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandle404(t *testing.T) {
	router := NewRouter()

	type args struct {
		method string
		target string
	}
	tests := []struct {
		name string
		args args
	}{
		{"GET /not-existing", args{"GET", "/not-existing"}},
		{"POST /other-not-existing", args{"POST", "/other-not-existing"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseWriter := httptest.NewRecorder()
			router.ServeHTTP(responseWriter, httptest.NewRequest(tt.args.method, tt.args.target, nil))

			if responseWriter.Code != http.StatusNotFound {
				t.Error("Did not get expected HTTP status code 404, got", responseWriter.Code)
			}
			if responseWriter.Body.String() != "{\"code\":404,\"message\":\"Not Found\"}\n" {
				t.Error("Did not get expected protocoll error, got", responseWriter.Body.String())
			}
		})
	}
}
