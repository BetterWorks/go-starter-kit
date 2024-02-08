package routes

import (
	"fmt"
	"net/http"

	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/http/jsonio"
	"github.com/go-chi/chi/v5"
)

// HealthRouter implements a router for healthcheck
func HealthRouter(r *chi.Mux, ns string) {
	prefix := fmt.Sprintf("/%s/health", ns)

	status := func(w http.ResponseWriter, r *http.Request) {
		data := models.Envelope{"meta": models.Envelope{"status": "healthy"}}
		jsonio.EncodeResponse(w, r, http.StatusOK, data)
	}

	r.Get(prefix, status)
}
