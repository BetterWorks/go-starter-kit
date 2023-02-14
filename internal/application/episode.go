package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/jasonsites/gosk-api/internal/types"
	"github.com/jasonsites/gosk-api/internal/validation"
)

type EpisodeServiceConfig struct {
	Logger *types.Logger    `validate:"required"`
	Repo   types.Repository `validate:"required"`
}

type episodeService struct {
	logger *types.Logger
	repo   types.Repository
}

func NewEpisodeService(c *EpisodeServiceConfig) (*episodeService, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	log := c.Logger.Log.With().Str("tags", "service,episode").Logger()
	logger := &types.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	service := &episodeService{
		logger: logger,
		repo:   c.Repo,
	}

	return service, nil
}

// Create
func (s *episodeService) Create(ctx context.Context, data any) (*types.JSONResponseSolo, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()

	result, err := s.repo.Create(ctx, data.(*types.EpisodeRequestData))
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	model := &types.Episode{}
	sr, err := model.SerializeResponse(result, true)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}
	res := sr.(*types.JSONResponseSolo)

	return res, nil
}

// Delete
func (s *episodeService) Delete(ctx context.Context, id uuid.UUID) error {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()

	if err := s.repo.Delete(ctx, id); err != nil {
		log.Error().Err(err).Msg("")
		return err
	}

	return nil
}

// Detail
func (s *episodeService) Detail(ctx context.Context, id uuid.UUID) (*types.JSONResponseSolo, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()

	result, err := s.repo.Detail(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	model := &types.Episode{}
	sr, err := model.SerializeResponse(result, true)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}
	res := sr.(*types.JSONResponseSolo)

	return res, nil
}

// List
func (s *episodeService) List(ctx context.Context, q types.QueryData) (*types.JSONResponseMult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()

	result, err := s.repo.List(ctx, q)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	model := &types.Episode{}
	sr, err := model.SerializeResponse(result, false)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}
	res := sr.(*types.JSONResponseMult)

	return res, nil
}

// Update
func (s *episodeService) Update(ctx context.Context, data any, id uuid.UUID) (*types.JSONResponseSolo, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()

	result, err := s.repo.Update(ctx, data.(*types.EpisodeRequestData), id)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	model := &types.Episode{}
	sr, err := model.SerializeResponse(result, true)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}
	res := sr.(*types.JSONResponseSolo)

	return res, nil
}
