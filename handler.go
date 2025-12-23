package httplusplus

import (
	"net/http"
)

type Handler func(ctx *Context) error
type Middleware func(next Handler) Handler

func Use(h Handler, middlewares ...Middleware) Handler {
	for i := range middlewares {
		h = middlewares[i](h)
	}
	return h
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
