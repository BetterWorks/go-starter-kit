package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health route implements an example route group for a specific domain resource
func Health(c Controller, ns string, r *gin.Engine) {
	prefix := "/" + ns + "/health"
	g := r.Group(prefix)

	status := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
	}

	g.GET("/", status)
}
