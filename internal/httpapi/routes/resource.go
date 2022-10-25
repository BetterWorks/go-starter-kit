package routes

import (
	"github.com/gin-gonic/gin"

	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
)

// Resource route implements an example route group for a specific domain resource
func Resource(c Controller, ns string, r *gin.Engine) {
	prefix := "/" + ns + "/resources"
	middleware := []gin.HandlerFunc{mw.LocalType("resource")}
	g := r.Group(prefix, middleware...)

	// temp := func(ctx *gin.Context) {
	// 	err := fmt.Errorf("some bad error happened")
	// 	ctx.AbortWithError(http.StatusInternalServerError, err)
	// }

	g.GET("/", c.List)
	g.GET("/:id", c.Detail)
	g.POST("/", c.Create)
	g.PATCH("/:id", c.Update)
	g.DELETE("/:id", c.Delete)
}
