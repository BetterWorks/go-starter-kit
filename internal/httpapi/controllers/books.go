package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jasonsites/gosk-api/internal/core/types"
)

// BookCreate
func (c *Controller) BookCreate(t string) gin.HandlerFunc {
	var data = &types.Book{}

	return func(ctx *gin.Context) {
		trace := ctx.MustGet("Trace").(types.Trace) // {Headers, RequestID}
		log := c.logger.Log.With().Str("req_id", trace.RequestID).Logger()
		log.Info().Msg("BookCreate Controller called")

		// TODO: validate body

		if err := ctx.BindJSON(data); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		// resource := c.application.Create(data)
		ctx.JSON(http.StatusCreated, data)
	}
}

// BookDelete
func (c *Controller) BookDelete(t string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := ctx.MustGet("Trace").(types.Trace) // {Headers, RequestID}
		log := c.logger.Log.With().Str("req_id", trace.RequestID).Logger()
		log.Info().Msg("BookDelete Controller called")

		id := ctx.Param("id")
		fmt.Printf("ID: %s", id)

		// c.application.Delete(id)
		ctx.Status(http.StatusNoContent)
	}
}

// BookDetail
func (c *Controller) BookDetail(t string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := ctx.MustGet("Trace").(types.Trace) // {Headers, RequestID}
		log := c.logger.Log.With().Str("req_id", trace.RequestID).Logger()
		log.Info().Msg("BookDetail Controller called")

		id := ctx.Param("id")
		fmt.Printf("ID: %s", id)

		// resource := c.application.Detail(id)
		// ctx.JSON(http.StatusOK, resource)

		ctx.Status(http.StatusOK)
	}
}

// BookList
func (c *Controller) BookList(t string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := ctx.MustGet("Trace").(types.Trace) // {Headers, RequestID}
		log := c.logger.Log.With().Str("req_id", trace.RequestID).Logger()
		log.Info().Msg("BookList Controller called")

		// TODO: get/bind/validate query from request
		// query := c.getQueryData(ctx, t)

		// resource := c.Application.List(query)
		// ctx.JSON(http.StatusOK, resource)

		ctx.Status(http.StatusOK)
	}
}

// BookUpdate
func (c *Controller) BookUpdate(t string) gin.HandlerFunc {
	var data = &types.Book{}

	return func(ctx *gin.Context) {
		trace := ctx.MustGet("Trace").(types.Trace) // {Headers, RequestID}
		log := c.logger.Log.With().Str("req_id", trace.RequestID).Logger()
		log.Info().Msg("BookUpdate Controller called")

		id := ctx.Param("id")
		fmt.Printf("ID: %s", id)

		// TODO: validate body

		if err := ctx.BindJSON(data); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		// resource := c.application.Update(data)
		// ctx.JSON(http.StatusOK, resource)

		ctx.Status(http.StatusOK)
	}
}
