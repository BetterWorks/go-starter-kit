package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jasonsites/gosk-api/internal/core/types"
	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
)

// BookCreate
func (c *Controller) BookCreate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Create Controller called")

		// TODO: validate body
		resource := &BookRequestBody{}
		if err := ctx.BodyParser(resource); err != nil {
			return err
		}

		model := types.Book(resource.Data.Properties)
		fmt.Printf("Model in Create Controller: %+v\n", model)

		result, err := c.application.Services.BookService.Create(&model)
		if err != nil {
			return err
		}
		fmt.Printf("Result in Create Controller: %+v\n", result)

		ctx.Status(http.StatusCreated)
		ctx.JSON(result)
		return nil
	}
}

// BookDelete
func (c *Controller) BookDelete() fiber.Handler {
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

// BookDetail
func (c *Controller) BookDetail() fiber.Handler {
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

// BookList
func (c *Controller) BookList() fiber.Handler {
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

// BookUpdate
func (c *Controller) BookUpdate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestID := ctx.Locals(mw.CorrelationContextKey).(*types.Trace).RequestID
		log := c.logger.Log.With().Str("req_id", requestID).Logger()
		log.Info().Msg("Update Controller called")

		id := ctx.Params("id")
		fmt.Printf("ID: %s", id)

		// TODO: validate body

		data := &BookRequestBody{}
		if err := ctx.BodyParser(data); err != nil {
			return err
		}
		// resource := c.application.Update(data)
		ctx.Status(http.StatusOK)
		ctx.JSON(fiber.Map{"data": "temp"})
		return nil
	}
}
