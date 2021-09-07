package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenApiDocumentation(t *testing.T) {
	dat, _ := ioutil.ReadFile("open-api-3.yaml")
	router := NewRouter()

	type args struct {
		method string
		target string
	}
	tests := []struct {
		name string
		args args
	}{
		{"GET /", args{"GET", "/"}},
		{"GET /api", args{"GET", "/api"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseWriter := httptest.NewRecorder()
			router.ServeHTTP(responseWriter, httptest.NewRequest(tt.args.method, tt.args.target, nil))

			if responseWriter.Code != http.StatusOK {
				t.Error("Did not get expected HTTP status code 200, got", responseWriter.Code)
			}
			if responseWriter.Body.String() != string(dat) {
				t.Error("Did not get expected open-api-3.yaml, got", responseWriter.Body.String())
			}
		})
	}
}
