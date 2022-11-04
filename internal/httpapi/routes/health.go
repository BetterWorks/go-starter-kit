package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// Health route implements an example route group for a specific domain resource
func HealthRouter(r *gin.Engine, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/health"
	g := r.Group(prefix)

	status := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
	}

	g.GET("/", status)
}
