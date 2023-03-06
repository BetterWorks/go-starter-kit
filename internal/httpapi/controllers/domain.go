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
func (c *Controller) Create(f func() *types.JSONRequestBody) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(types.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Create Controller called")

		resource := f()
		if err := ctx.BodyParser(resource); err != nil {
			message := "error parsing request body"
			log.Error().Err(err).Msg(message)
			return fiber.NewError(http.StatusBadRequest, message)
		}

		// TODO: validation errors bypass default error handler
		if err := validateBody(resource, log); err != nil {
			log.Error().Msg("validation error")
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(err.Errors)
			return nil
		}

		model := resource.Data.Properties
		result, err := c.service.Create(ctx.Context(), model)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(result)
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
			log.Error().Err(err).Send()
			return err
		}

		if err := c.service.Delete(ctx.Context(), uuid); err != nil {
			log.Error().Err(err).Send()
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
			log.Error().Err(err).Send()
			return err
		}

		result, err := c.service.Detail(ctx.Context(), uuid)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(result)
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
			log.Error().Err(err).Send()
			return err
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(result)
	}
}

// Update
func (c *Controller) Update(f func() *types.JSONRequestBody) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(types.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Update Controller called")

		idString := ctx.Params("id")
		id, err := uuid.Parse(idString)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}

		resource := f()
		if err := ctx.BodyParser(resource); err != nil {
			log.Error().Err(err).Send()
			return err
		}

		if err := validateBody(resource, log); err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(err)
			return nil
		}

		model := resource.Data.Properties // TODO: problem here with ID
		result, err := c.service.Update(ctx.Context(), model, id)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(result)
	}
}
