package domain

import (
	"context"
	"fmt"

	"github.com/BetterWorks/gosk-api/internal/core/interfaces"
	"github.com/BetterWorks/gosk-api/internal/core/logger"
	"github.com/BetterWorks/gosk-api/internal/core/models"
	"github.com/BetterWorks/gosk-api/internal/core/query"
	"github.com/BetterWorks/gosk-api/internal/core/trace"
	"github.com/BetterWorks/gosk-api/internal/core/validation"
	"github.com/google/uuid"
)

// ExampleServiceConfig defines the input to NewExampleService
type ExampleServiceConfig struct {
	Logger *logger.Logger               `validate:"required"`
	Repo   interfaces.ExampleRepository `validate:"required"`
}

// exampleService
type exampleService struct {
	logger *logger.Logger
	repo   interfaces.ExampleRepository
}

// NewExampleService returns a new exampleService instance
func NewExampleService(c *ExampleServiceConfig) (*exampleService, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	log := c.Logger.Log.With().Str("tags", "service,example").Logger()
	logger := &logger.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	service := &exampleService{
		logger: logger,
		repo:   c.Repo,
	}

	return service, nil
}

// Create
func (s *exampleService) Create(ctx context.Context, data any) (interfaces.DomainModel, error) {
	traceID := trace.GetTraceIDFromContext(ctx)
	log := s.logger.CreateContextLogger(traceID)

	d, ok := data.(*models.ExampleInputData)
	if !ok {
		err := fmt.Errorf("example input data assertion error")
		log.Error().Err(err).Send()
		return nil, err
	}

	model, err := s.repo.Create(ctx, d)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return model, nil
}

// Delete
func (s *exampleService) Delete(ctx context.Context, id uuid.UUID) error {
	traceID := trace.GetTraceIDFromContext(ctx)
	log := s.logger.CreateContextLogger(traceID)

	if err := s.repo.Delete(ctx, id); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}

// Detail
func (s *exampleService) Detail(ctx context.Context, id uuid.UUID) (interfaces.DomainModel, error) {
	traceID := trace.GetTraceIDFromContext(ctx)
	log := s.logger.CreateContextLogger(traceID)

	model, err := s.repo.Detail(ctx, id)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return model, nil
}

// List
func (s *exampleService) List(ctx context.Context, q query.QueryData) (interfaces.DomainModel, error) {
	traceID := trace.GetTraceIDFromContext(ctx)
	log := s.logger.CreateContextLogger(traceID)

	model, err := s.repo.List(ctx, q)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return model, nil
}

// Update
func (s *exampleService) Update(ctx context.Context, data any, id uuid.UUID) (interfaces.DomainModel, error) {
	traceID := trace.GetTraceIDFromContext(ctx)
	log := s.logger.CreateContextLogger(traceID)

	d, ok := data.(*models.ExampleInputData)
	if !ok {
		err := fmt.Errorf("example input data assertion error")
		log.Error().Err(err).Send()
		return nil, err
	}

	model, err := s.repo.Update(ctx, d, id)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return model, nil
}
