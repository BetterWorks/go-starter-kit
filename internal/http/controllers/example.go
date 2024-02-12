package controllers

import (
	"net/http"

	"github.com/BetterWorks/go-starter-kit/internal/core/app"
	"github.com/BetterWorks/go-starter-kit/internal/core/cerror"
	"github.com/BetterWorks/go-starter-kit/internal/core/interfaces"
	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/core/trace"
	"github.com/BetterWorks/go-starter-kit/internal/http/jsonio"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// ExampleControllerConfig defines the input to NewExampleController
type ExampleControllerConfig struct {
	Logger      *logger.CustomLogger
	QueryConfig *QueryConfig
	Service     interfaces.ExampleService
}

// ExampleController
type ExampleController struct {
	logger  *logger.CustomLogger
	query   *queryHandler
	service interfaces.ExampleService
}

// NewExampleController returns a new ExampleController instance
func NewExampleController(c *ExampleControllerConfig) (*ExampleController, error) {
	if err := app.Validator.Validate.Struct(c); err != nil {
		return nil, err
	}

	queryHandler, err := NewQueryHandler(c.QueryConfig)
	if err != nil {
		return nil, err
	}
	ctrl := &ExampleController{
		logger:  c.Logger,
		query:   queryHandler,
		service: c.Service,
	}

	return ctrl, nil
}

// Create
func (c *ExampleController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := trace.GetTraceIDFromContext(ctx)
		log := c.logger.CreateContextLogger(traceID)

		body := &models.ExampleRequest{
			Data: &models.ExampleRequestResource{
				Attributes: models.ExampleRequestAttributes{},
			},
		}

		if err := jsonio.DecodeRequest(w, r, body); err != nil {
			err = cerror.NewValidationError(err, "request body decode error")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		data := &body.Data.Attributes
		if err := data.Validate(); err != nil {
			err = cerror.NewValidationError(err, "invalid request body")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		model, err := c.service.Create(ctx, data)
		if err != nil {
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		response, err := model.FormatResponse()
		if err != nil {
			err = cerror.NewInternalServerError(err, "model format response error")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		jsonio.EncodeResponse(w, r, http.StatusCreated, response)
	}
}

// Delete
func (c *ExampleController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := trace.GetTraceIDFromContext(ctx)
		log := c.logger.CreateContextLogger(traceID)

		id := chi.URLParam(r, "id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			err = cerror.NewInternalServerError(err, "error parsing resource id")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		if err := c.service.Delete(ctx, uuid); err != nil {
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// Detail
func (c *ExampleController) Detail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := trace.GetTraceIDFromContext(ctx)
		log := c.logger.CreateContextLogger(traceID)

		id := chi.URLParam(r, "id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			err = cerror.NewInternalServerError(err, "error parsing resource id")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		model, err := c.service.Detail(ctx, uuid)
		if err != nil {
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		response, err := model.FormatResponse()
		if err != nil {
			err = cerror.NewInternalServerError(err, "error formatting response from model")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		jsonio.EncodeResponse(w, r, http.StatusOK, response)
	}
}

// List
func (c *ExampleController) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := trace.GetTraceIDFromContext(ctx)
		log := c.logger.CreateContextLogger(traceID)

		qs := []byte(r.URL.RawQuery)
		query := c.query.parseQuery(qs)

		model, err := c.service.List(ctx, *query)
		if err != nil {
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		response, err := model.FormatResponse()
		if err != nil {
			err = cerror.NewInternalServerError(err, "error formatting response from model")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		jsonio.EncodeResponse(w, r, http.StatusOK, response)
	}
}

// Update
func (c *ExampleController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := trace.GetTraceIDFromContext(ctx)
		log := c.logger.CreateContextLogger(traceID)

		id := chi.URLParam(r, "id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			err = cerror.NewValidationError(err, "resource id parse error")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		body := &models.ExampleRequest{
			Data: &models.ExampleRequestResource{
				Attributes: models.ExampleRequestAttributes{},
			},
		}

		if err := jsonio.DecodeRequest(w, r, body); err != nil {
			err = cerror.NewValidationError(err, "request body decode error")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		data := &body.Data.Attributes
		if err := data.Validate(); err != nil {
			err = cerror.NewValidationError(err, "invalid request body")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		model, err := c.service.Update(ctx, data, uuid)
		if err != nil {
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		response, err := model.FormatResponse()
		if err != nil {
			err = cerror.NewInternalServerError(err, "model format response error")
			log.Error(err.Error())
			jsonio.EncodeError(w, r, err)
			return
		}

		jsonio.EncodeResponse(w, r, http.StatusOK, response)
	}
}
