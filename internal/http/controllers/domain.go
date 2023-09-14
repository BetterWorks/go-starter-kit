package controllers

import (
	"fmt"
	"net/http"

	"github.com/BetterWorks/gosk-api/internal/core/cerror"
	"github.com/BetterWorks/gosk-api/internal/core/interfaces"
	"github.com/BetterWorks/gosk-api/internal/core/logger"
	"github.com/BetterWorks/gosk-api/internal/core/trace"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Config defines the input to NewController
type Config struct {
	Logger      *logger.Logger
	QueryConfig *QueryConfig
	Service     interfaces.Service
}

// Controller
type Controller struct {
	logger  *logger.Logger
	query   *queryHandler
	service interfaces.Service
}

// NewController returns a new Controller instance
func NewController(c *Config) *Controller {
	queryHandler := NewQueryHandler(c.QueryConfig)
	return &Controller{
		logger:  c.Logger,
		query:   queryHandler,
		service: c.Service,
	}
}

// Create
func (c *Controller) Create(f func() *RequestBody) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := trace.GetTraceIDFromContext(ctx)
		log := c.logger.CreateContextLogger(traceID)

		resource := f()
		if err := c.JSONDecode(w, r, resource); err != nil {
			err = cerror.NewInternalServerError("request body decode error")
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		// TODO: validation errors bypass default error handler (rethink this)
		if response := validateBody(resource, log); response != nil {
			err := fmt.Errorf("validation error(s) %+v", response)
			log.Error().Err(err).Send()
			c.JSONEncode(w, r, http.StatusBadRequest, response)
			return
		}

		data := resource.Data.Attributes
		model, err := c.service.Create(ctx, data)
		if err != nil {
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		response, err := model.FormatResponse()
		if err != nil {
			err = cerror.NewInternalServerError("model format response error")
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		c.JSONEncode(w, r, http.StatusCreated, response)
	}
}

// Delete
func (c *Controller) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := trace.GetTraceIDFromContext(ctx)
		log := c.logger.CreateContextLogger(traceID)

		id := chi.URLParam(r, "id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			err = cerror.NewInternalServerError("error parsing resource id")
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		if err := c.service.Delete(ctx, uuid); err != nil {
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// Detail
func (c *Controller) Detail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := trace.GetTraceIDFromContext(ctx)
		log := c.logger.CreateContextLogger(traceID)

		id := chi.URLParam(r, "id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			err = cerror.NewInternalServerError("error parsing resource id")
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		model, err := c.service.Detail(ctx, uuid)
		if err != nil {
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		response, err := model.FormatResponse()
		if err != nil {
			err = cerror.NewInternalServerError("error formatting response from model")
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		c.JSONEncode(w, r, http.StatusOK, response)
	}
}

// List
func (c *Controller) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := trace.GetTraceIDFromContext(ctx)
		log := c.logger.CreateContextLogger(traceID)

		qs := []byte(r.URL.RawQuery)
		query := c.query.parseQuery(qs)

		model, err := c.service.List(ctx, *query)
		if err != nil {
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		response, err := model.FormatResponse()
		if err != nil {
			err = cerror.NewInternalServerError("error formatting response from model")
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		c.JSONEncode(w, r, http.StatusOK, response)
	}
}

// Update
func (c *Controller) Update(f func() *RequestBody) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceID := trace.GetTraceIDFromContext(ctx)
		log := c.logger.CreateContextLogger(traceID)

		id := chi.URLParam(r, "id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			err = cerror.NewInternalServerError("resource id parse error")
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		resource := f()
		if err := c.JSONDecode(w, r, resource); err != nil {
			err = cerror.NewInternalServerError("request body decode error")
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		// TODO: validation errors bypass default error handler
		if response := validateBody(resource, log); response != nil {
			err := fmt.Errorf("validation error(s) %+v", response)
			log.Error().Err(err).Send()
			c.JSONEncode(w, r, http.StatusBadRequest, response)
			return
		}

		data := resource.Data.Attributes // TODO: problem here with ID
		model, err := c.service.Update(ctx, data, uuid)
		if err != nil {
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		response, err := model.FormatResponse()
		if err != nil {
			err = cerror.NewInternalServerError("model format response error")
			log.Error().Err(err).Send()
			c.Error(w, r, err)
			return
		}

		c.JSONEncode(w, r, http.StatusCreated, response)
	}
}
