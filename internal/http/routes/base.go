package routes

import (
	"fmt"
	"net/http"

	ctrl "github.com/BetterWorks/gosk-api/internal/http/controllers"
	"github.com/go-chi/chi/v5"
)

// BaseRouter only exists to easily verify a working app and should normally be removed
func BaseRouter(r *chi.Mux, c *ctrl.Controller, ns string) {
	prefix := fmt.Sprintf("/%s", ns)

	get := func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		host := r.Host
		path := r.URL.Path
		remoteAddress := r.RemoteAddr

		data := ctrl.Envelope{
			"data": "base router is working...",
			"request": ctrl.Envelope{
				"headers":       headers,
				"host":          host,
				"path":          path,
				"remoteAddress": remoteAddress,
			},
		}

		c.JSONEncode(w, r, http.StatusOK, data)
	}

	r.Route(prefix, func(r chi.Router) {
		r.Get("/", get)
	})
}
