package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/jasonsites/gosk-api/internal/core/types"
	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// Resource2Router implements an example router group for a specific domain resource
func BookRouter(r *gin.Engine, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/books"
	t := types.ResourceType.Book
	g := r.Group(prefix)

	// g.GET("/", c.BookList(t))
	// g.GET("/:id", c.BookDetail(t))
	g.POST("/", c.BookCreate(t))
	// g.PATCH("/:id", c.BookUpdate(t))
	// g.DELETE("/:id", c.BookDelete(t))
}
