package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// Base route only exists to easily verify a working app and should normally be removed
func BaseRouter(r *gin.Engine, c *ctrl.Controller, ns string) {
	prefix := "/" + ns
	g := r.Group(prefix)

	get := func(ctx *gin.Context) {
		cookies := ctx.Request.Cookies()
		headers := ctx.Request.Header
		host := ctx.Request.Host
		remoteAddress := ctx.Request.RemoteAddr
		requestURI := ctx.Request.RequestURI
		url := ctx.Request.URL.String()

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"data": "base router is working...",
			"request": gin.H{
				"cookies":       cookies,
				"headers":       headers,
				"host":          host,
				"remoteAddress": remoteAddress,
				"requestURI":    requestURI,
				"url":           url,
			},
		})
	}

	g.GET("/", get)
}
