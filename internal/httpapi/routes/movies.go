package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/jasonsites/gosk-api/internal/core/types"
	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// Resource route implements an example router group for a specific domain resource
func MovieRouter(r *gin.Engine, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/movies"
	t := types.ResourceType.Movie
	// middlewares := []gin.HandlerFunc{mw.LocalType(t)}
	g := r.Group(prefix)

	// temp := func(ctx *gin.Context) {
	// 	err := fmt.Errorf("some bad error happened")
	// 	ctx.AbortWithError(http.StatusInternalServerError, err)
	// }

	// g.GET("/", c.MovieList(t))
	// g.GET("/:id", c.MovieDetail(t))
	g.POST("/", c.MovieCreate(t))
	// g.PATCH("/:id", c.MovieUpdate(t))
	// g.DELETE("/:id", c.MovieDelete(t))
}
