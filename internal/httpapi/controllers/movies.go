package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jasonsites/gosk-api/internal/core/types"
	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
)

// MovieCreate
func (c *Controller) MovieCreate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Create Controller called")

		// TODO: validate body
		data := &MovieRequestBody{}
		if err := ctx.BodyParser(data); err != nil {
			return err
		}
		fmt.Printf("DATA in MovieCreate Controller: %+v\n", data)

		// resource := types.Book(data.Data.Properties)
		// resource := c.application.Create(data)
		ctx.Status(http.StatusCreated)
		ctx.JSON(data)
		return nil
	}
}

// MovieDelete
func (c *Controller) MovieDelete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("MovieDelete Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		// c.application.Delete(id)
		ctx.Status(http.StatusNoContent)
		return nil
	}
}

// MovieDetail
func (c *Controller) MovieDetail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("MovieDetail Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		// resource := c.application.Detail(id)

		ctx.Status(http.StatusOK)
		ctx.JSON(fiber.Map{"data": "temp"})
		return nil
	}
}

// MovieList
func (c *Controller) MovieList() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("MovieList Controller called")

		// TODO: get/bind/validate query from request
		// query := c.getQueryData(ctx, t)

		// resource := c.Application.List(query)

		ctx.Status(http.StatusOK)
		ctx.JSON(fiber.Map{"data": "temp"})
		return nil
	}
}

// MovieUpdate
func (c *Controller) MovieUpdate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Update Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		// TODO: validate body
		data := &MovieRequestBody{}
		if err := ctx.BodyParser(data); err != nil {
			return err
		}
		fmt.Printf("DATA in MovieUpdate Controller: %+v\n", data)

		// resource := c.application.Update(data)
		ctx.Status(http.StatusOK)
		ctx.JSON(fiber.Map{"data": "temp"})
		return nil
	}
}
