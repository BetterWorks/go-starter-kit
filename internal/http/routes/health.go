package routes

import (
	"fmt"
	"net/http"

	ctrl "github.com/BetterWorks/gosk-api/internal/http/controllers"
	"github.com/go-chi/chi/v5"
)

// HealthRouter implements a router for healthcheck
func HealthRouter(r *chi.Mux, c *ctrl.Controller, ns string) {
	prefix := fmt.Sprintf("/%s/health", ns)

	status := func(w http.ResponseWriter, r *http.Request) {
		data := ctrl.Envelope{"meta": ctrl.Envelope{"status": "healthy"}}
		c.JSONEncode(w, r, http.StatusOK, data)
	}

	r.Get(prefix, status)
}
