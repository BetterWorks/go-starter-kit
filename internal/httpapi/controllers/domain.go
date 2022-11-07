package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jasonsites/gosk-api/internal/core/types"
	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
)

type Config struct {
	Application types.Application
	Logger      *types.Logger
}

type Controller struct {
	application types.Application
	logger      *types.Logger
}

func NewController(c *Config) *Controller {
	return &Controller{
		application: c.Application,
		logger:      c.Logger,
	}
}

// Create
func (c *Controller) Create(data any, t string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Create Controller called")

		// TODO: validate body

		if err := ctx.BodyParser(data); err != nil {
			return err
		}
		// resource := c.application.Create(data)
		ctx.Status(http.StatusCreated)
		ctx.JSON(data)
		return nil
	}
}

// Delete
func (c *Controller) Delete(t string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Delete Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		// c.application.Delete(id)
		ctx.Status(http.StatusNoContent)
		return nil
	}
}

// Detail
func (c *Controller) Detail(t string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Detail Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		// resource := c.application.Detail(id)

		ctx.Status(http.StatusOK)
		ctx.JSON(fiber.Map{"data": "temp"})
		return nil
	}
}

// List
func (c *Controller) List(t string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("List Controller called")

		// TODO: get/bind/validate query from request
		// query := c.getQueryData(ctx, t)

		// resource := c.Application.List(query)

		ctx.Status(http.StatusOK)
		ctx.JSON(fiber.Map{"data": "temp"})
		return nil
	}
}

// Update
func (c *Controller) Update(data any, t string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Update Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		// TODO: validate body

		if err := ctx.BodyParser(data); err != nil {
			return err
		}
		// resource := c.application.Update(data)
		ctx.Status(http.StatusOK)
		ctx.JSON(fiber.Map{"data": "temp"})
		return nil
	}
}
