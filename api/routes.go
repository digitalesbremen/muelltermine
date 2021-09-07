package api

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "Open Api documentation (yaml)",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: OpenApiDocumentation,
	},
	Route{
		Name:        "Open Api documentation (yaml)",
		Method:      "GET",
		Pattern:     "/api",
		HandlerFunc: OpenApiDocumentation,
	},
}
