package httplusplus

import (
	"net/http"
)

type Route struct {
	path    string
	handler Handler
}

type AppRouter struct {
	Mux *http.ServeMux
}

func NewRouter() *AppRouter {
	return &AppRouter{
		Mux: http.NewServeMux(),
	}
}

func (a *AppRouter) Route(
	path string,
	handler Handler,
) *AppRouter {
	a.Mux.HandleFunc(
		path,
		handler.toStdHandler(),
	)
	return a
}
