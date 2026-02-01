package httplusplus

import (
	"net/http"
)

type Handler func(ctx *Context) error
type Middleware func(next Handler) Handler

func (r *Route) Use(middlewares ...Middleware) *Route {
	for i := len(middlewares) - 1; i >= 0; i-- {
		r.handler = middlewares[i](r.handler)
	}

	return r
}

func (handler Handler) toStdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{ResponseWriter: w, Request: r}
		if err := handler(ctx); err != nil {
			if fError, ok := err.(FrameworkError); ok {
				ctx.Json(fError.Status, fError)
				return
			}
			ctx.Json(
				http.StatusInternalServerError,
				map[string]string{"message": "Internal server error"},
			)
		}
	}
}

// Nest a different ServeMux prefix path must be in the pattern, "/<prefix>"
func (a *AppRouter) NestHandler(prefix string, app *AppRouter) *AppRouter {
	a.Mux.Handle(prefix+"/", http.StripPrefix(prefix, app.Mux))
	return a
}
