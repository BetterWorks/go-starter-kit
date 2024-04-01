package domain

import (
	"context"

	"github.com/BetterWorks/go-starter-kit/internal/core/app"
	"github.com/BetterWorks/go-starter-kit/internal/core/interfaces"
	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/core/query"
	"github.com/BetterWorks/go-starter-kit/internal/core/trace"
	"github.com/google/uuid"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// ExampleServiceConfig defines the input to NewExampleService
type ExampleServiceConfig struct {
	Logger         *logger.CustomLogger         `validate:"required"`
	Repo           interfaces.ExampleRepository `validate:"required"`
	NewRelicClient *newrelic.Application
}

// exampleService
type exampleService struct {
	logger         *logger.CustomLogger
	repo           interfaces.ExampleRepository
	newRelicClient *newrelic.Application
}

// NewExampleService returns a new exampleService instance
func NewExampleService(c *ExampleServiceConfig) (*exampleService, error) {
	if err := app.Validator.Validate.Struct(c); err != nil {
		return nil, err
	}

	service := &exampleService{
		logger:         c.Logger,
		repo:           c.Repo,
		newRelicClient: c.NewRelicClient,
	}

	return service, nil
}

// Create
func (s *exampleService) Create(ctx context.Context, data *models.ExampleRequestAttributes) (*models.ExampleDomainModel, error) {
	txn := s.newRelicClient.StartTransaction("exampleService.Create")
	defer txn.End()

	traceID := trace.GetTraceIDFromContext(ctx)
	log := s.logger.CreateContextLogger(traceID)

	seg := txn.StartSegment("exampleRepository.Create")
	model, err := s.repo.Create(ctx, data)
	seg.End()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return model, nil
}

// Delete
func (s *exampleService) Delete(ctx context.Context, id uuid.UUID) error {
	txn := s.newRelicClient.StartTransaction("exampleService.Delete")
	defer txn.End()

	traceID := trace.GetTraceIDFromContext(ctx)
	log := s.logger.CreateContextLogger(traceID)

	seg := txn.StartSegment("exampleRepository.Delete")
	err := s.repo.Delete(ctx, id)
	seg.End()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

// Detail
func (s *exampleService) Detail(ctx context.Context, id uuid.UUID) (*models.ExampleDomainModel, error) {
	txn := s.newRelicClient.StartTransaction("exampleService.Detail")
	defer txn.End()

	traceID := trace.GetTraceIDFromContext(ctx)
	log := s.logger.CreateContextLogger(traceID)

	seg := txn.StartSegment("exampleRepository.Detail")
	model, err := s.repo.Detail(ctx, id)
	seg.End()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return model, nil
}

// List
func (s *exampleService) List(ctx context.Context, q query.QueryData) (*models.ExampleDomainModel, error) {
	txn := s.newRelicClient.StartTransaction("exampleService.List")
	defer txn.End()

	traceID := trace.GetTraceIDFromContext(ctx)
	log := s.logger.CreateContextLogger(traceID)

	seg := txn.StartSegment("exampleRepository.List")
	model, err := s.repo.List(ctx, q)
	seg.End()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return model, nil
}

// Update
func (s *exampleService) Update(ctx context.Context, data *models.ExampleRequestAttributes, id uuid.UUID) (*models.ExampleDomainModel, error) {
	txn := s.newRelicClient.StartTransaction("exampleService.Update")
	defer txn.End()

	traceID := trace.GetTraceIDFromContext(ctx)
	log := s.logger.CreateContextLogger(traceID)

	seg := txn.StartSegment("exampleRepository.Update")
	model, err := s.repo.Update(ctx, data, id)
	seg.End()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return model, nil
}
