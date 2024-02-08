package middleware

import (
	"net/http"

	"github.com/BetterWorks/go-starter-kit/internal/core/cerror"
	"github.com/BetterWorks/go-starter-kit/internal/http/jsonio"
	"github.com/go-chi/chi/v5"
)

// NotFound
func NotFound(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		tctx := chi.NewRouteContext()
		if !rctx.Routes.Match(tctx, r.Method, r.URL.Path) {
			err := cerror.NewNotFoundError(nil, "path not found")
			jsonio.EncodeError(w, r, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
