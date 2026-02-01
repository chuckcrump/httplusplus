package httplusplus

import (
	"net/http"
)

type Route struct {
	router  *AppRouter
	path    string
	handler Handler
}

type AppRouter struct {
	Mux    *http.ServeMux
	routes []*Route
}

func NewRouter() *AppRouter {
	return &AppRouter{
		Mux: http.NewServeMux(),
	}
}

func (a *AppRouter) Route(
	path string,
	handler Handler,
) *Route {
	r := &Route{
		router:  a,
		path:    path,
		handler: handler,
	}
	a.routes = append(a.routes, r)
	return r
}
