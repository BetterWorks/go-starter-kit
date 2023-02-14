package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jasonsites/gosk-api/internal/types"
)

// Config
type Config struct {
	Service types.Service
	Logger  *types.Logger
}

// Controller
type Controller struct {
	service types.Service
	logger  *types.Logger
}

// NewController
func NewController(c *Config) *Controller {
	return &Controller{
		service: c.Service,
		logger:  c.Logger,
	}
}

// Create
func (c *Controller) Create(f func() *JSONRequestBody) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(types.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Create Controller called")

		resource := f()
		if err := ctx.BodyParser(resource); err != nil {
			log.Error().Err(err).Msg("")
			return err
		}

		model := resource.Data.Properties
		result, err := c.service.Create(ctx.Context(), model)
		if err != nil {
			log.Error().Err(err).Msg("")
			return err
		}

		ctx.Status(http.StatusCreated)
		ctx.JSON(result)
		return nil
	}
}

// Delete
func (c *Controller) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(types.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()

		id := ctx.Params("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			log.Error().Err(err).Msg("")
			return err
		}

		if err := c.service.Delete(ctx.Context(), uuid); err != nil {
			log.Error().Err(err).Msg("")
			return err
		}
		ctx.Status(http.StatusNoContent)
		return nil
	}
}

// Detail
func (c *Controller) Detail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(types.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Detail Controller called")

		id := ctx.Params("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			log.Error().Err(err).Msg("")
			return err
		}

		result, err := c.service.Detail(ctx.Context(), uuid)
		if err != nil {
			log.Error().Err(err).Msg("")
			return err
		}

		ctx.Status(http.StatusOK)
		ctx.JSON(result)
		return nil
	}
}

// List
func (c *Controller) List() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(types.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("List Controller called")

		qs := ctx.Request().URI().QueryString()
		query := parseQuery(qs)

		result, err := c.service.List(ctx.Context(), *query)
		if err != nil {
			log.Error().Err(err).Msg("")
			return err
		}

		ctx.Status(http.StatusOK)
		ctx.JSON(result)
		return nil
	}
}

// Update
func (c *Controller) Update(f func() *JSONRequestBody) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(types.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Update Controller called")

		idString := ctx.Params("id")
		id, err := uuid.Parse(idString)
		if err != nil {
			log.Error().Err(err).Msg("")
			return err
		}

		// TODO: validate body

		resource := f()
		if err := ctx.BodyParser(resource); err != nil {
			log.Error().Err(err).Msg("")
			return err
		}

		model := resource.Data.Properties // TODO: problem here with ID
		result, err := c.service.Update(ctx.Context(), model, id)
		if err != nil {
			log.Error().Err(err).Msg("")
			return err
		}

		ctx.Status(http.StatusOK)
		ctx.JSON(result)
		return nil
	}
}
