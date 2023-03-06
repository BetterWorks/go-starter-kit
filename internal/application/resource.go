package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/jasonsites/gosk-api/internal/types"
	"github.com/jasonsites/gosk-api/internal/validation"
)

type ResourceServiceConfig struct {
	Logger *types.Logger    `validate:"required"`
	Repo   types.Repository `validate:"required"`
}

type resourceService struct {
	logger *types.Logger
	repo   types.Repository
}

func NewResourceService(c *ResourceServiceConfig) (*resourceService, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	log := c.Logger.Log.With().Str("tags", "service,resource").Logger()
	logger := &types.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	service := &resourceService{
		logger: logger,
		repo:   c.Repo,
	}

	return service, nil
}

// Create
func (s *resourceService) Create(ctx context.Context, data any) (*types.JSONResponseSolo, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()

	result, err := s.repo.Create(ctx, data.(*types.ResourceRequestData))
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	model := &types.Resource{}
	sr, err := model.SerializeResponse(result, true)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	res := sr.(*types.JSONResponseSolo)

	return res, nil
}

// Delete
func (s *resourceService) Delete(ctx context.Context, id uuid.UUID) error {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()

	if err := s.repo.Delete(ctx, id); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}

// Detail
func (s *resourceService) Detail(ctx context.Context, id uuid.UUID) (*types.JSONResponseSolo, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()

	result, err := s.repo.Detail(ctx, id)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	model := &types.Resource{}
	sr, err := model.SerializeResponse(result, true)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	res := sr.(*types.JSONResponseSolo)

	return res, nil
}

// List
func (s *resourceService) List(ctx context.Context, q types.QueryData) (*types.JSONResponseMult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()

	result, err := s.repo.List(ctx, q)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	model := &types.Resource{}
	sr, err := model.SerializeResponse(result, false)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	res := sr.(*types.JSONResponseMult)

	return res, nil
}

// Update
func (s *resourceService) Update(ctx context.Context, data any, id uuid.UUID) (*types.JSONResponseSolo, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()

	result, err := s.repo.Update(ctx, data.(*types.ResourceRequestData), id)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	model := &types.Resource{}
	sr, err := model.SerializeResponse(result, true)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	res := sr.(*types.JSONResponseSolo)

	return res, nil
}
