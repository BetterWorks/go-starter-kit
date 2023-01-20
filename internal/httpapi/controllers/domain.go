package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
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
		fmt.Printf("\n\nINSIDE CREATE HANDLER\n\n")
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		fmt.Printf("REQUEST ID: %+v\n", requestID)
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Create Controller called")

		resource := f()

		if err := ctx.BodyParser(resource); err != nil {
			fmt.Printf("Error in BodyParser %+v\n", err)
			return err
		}
		fmt.Printf("Resource in Create Controller: %+v\n", *resource)

		model := resource.Data.Properties
		fmt.Printf("Model in Create Controller: %+v\n", model)

		result, err := c.service.Create(model)
		if err != nil {
			return err
		}
		fmt.Printf("Result in Create Controller: %+v\n", result)

		ctx.Status(http.StatusCreated)
		ctx.JSON(result)
		return nil
	}
}

// Delete
func (c *Controller) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Delete Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		uuid, err := uuid.Parse(id)
		if err != nil {
			return err
		}

		if err := c.service.Delete(uuid); err != nil {
			return err
		}
		ctx.Status(http.StatusNoContent)
		return nil
	}
}

// Detail
func (c *Controller) Detail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Detail Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID in Controller: %s\n", id)

		uuid, err := uuid.Parse(id)
		if err != nil {
			return err
		}

		result, err := c.service.Detail(uuid)
		if err != nil {
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
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("List Controller called")

		// TODO: get/bind/validate query from request

		query := &types.ListMeta{} // c.getQueryData(ctx)
		result, err := c.service.List(query)
		if err != nil {
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
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Update Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		// TODO: validate body

		resource := f()

		if err := ctx.BodyParser(resource); err != nil {
			fmt.Printf("Error in BodyParser %+v\n", err)
			return err
		}
		fmt.Printf("Resource in Update Controller: %+v\n", *resource)

		model := resource.Data.Properties // TODO: problem here with ID
		fmt.Printf("Model in Update Controller: %+v\n", model)

		result, err := c.service.Update(model)
		if err != nil {
			return err
		}
		fmt.Printf("Result in Update Controller: %+v\n", result)

		ctx.Status(http.StatusOK)
		ctx.JSON(result)
		return nil
	}
}
