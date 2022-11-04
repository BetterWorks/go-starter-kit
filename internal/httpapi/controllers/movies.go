package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jasonsites/gosk-api/internal/core/types"
)

// MovieCreate
func (c *Controller) MovieCreate(t string) gin.HandlerFunc {
	var data = &types.Movie{}

	return func(ctx *gin.Context) {
		trace := ctx.MustGet("Trace").(types.Trace) // {Headers, RequestID}
		log := c.logger.Log.With().Str("req_id", trace.RequestID).Logger()
		log.Info().Msg("MovieCreate Controller called")

		// TODO: validate body

		if err := ctx.BindJSON(data); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		// resource := c.application.Create(data)
		ctx.JSON(http.StatusCreated, data)
	}
}

// MovieDelete
func (c *Controller) MovieDelete(t string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := ctx.MustGet("Trace").(types.Trace) // {Headers, RequestID}
		log := c.logger.Log.With().Str("req_id", trace.RequestID).Logger()
		log.Info().Msg("MovieDelete Controller called")

		id := ctx.Param("id")
		fmt.Printf("ID: %s", id)

		// c.application.Delete(id)
		ctx.Status(http.StatusNoContent)
	}
}

// MovieDetail
func (c *Controller) MovieDetail(t string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := ctx.MustGet("Trace").(types.Trace) // {Headers, RequestID}
		log := c.logger.Log.With().Str("req_id", trace.RequestID).Logger()
		log.Info().Msg("MovieDetail Controller called")

		id := ctx.Param("id")
		fmt.Printf("ID: %s", id)

		// resource := c.application.Detail(id)
		// ctx.JSON(http.StatusOK, resource)

		ctx.Status(http.StatusOK)
	}
}

// MovieList
func (c *Controller) MovieList(t string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trace := ctx.MustGet("Trace").(types.Trace) // {Headers, RequestID}
		log := c.logger.Log.With().Str("req_id", trace.RequestID).Logger()
		log.Info().Msg("MovieList Controller called")

		// TODO: get/bind/validate query from request
		// query := c.getQueryData(ctx, t)

		// resource := c.Application.List(query)
		// ctx.JSON(http.StatusOK, resource)

		ctx.Status(http.StatusOK)
	}
}

// MovieUpdate
func (c *Controller) MovieUpdate(t string) gin.HandlerFunc {
	var data = &types.Movie{}

	return func(ctx *gin.Context) {
		trace := ctx.MustGet("Trace").(types.Trace) // {Headers, RequestID}
		log := c.logger.Log.With().Str("req_id", trace.RequestID).Logger()
		log.Info().Msg("MovieUpdate Controller called")

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
