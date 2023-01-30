package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/jasonsites/gosk-api/internal/types"
)

type EpisodeServiceConfig struct {
	Logger *types.Logger    `validate:"required"`
	Repo   types.Repository `validate:"required"`
}

type episodeService struct {
	logger *types.Logger
	repo   types.Repository
}

func NewEpisodeService(c *EpisodeServiceConfig) *episodeService {
	log := c.Logger.Log.With().Str("tags", "service,episode").Logger()
	logger := &types.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	return &episodeService{
		logger: logger,
		repo:   c.Repo,
	}
}

// Create
func (s *episodeService) Create(ctx context.Context, data any) (*types.JSONResponseSolo, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Episode Service Create called")

	result, err := s.repo.Create(ctx, data.(*types.EpisodeRequestData))
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	model := &types.Episode{}
	sr, err := model.SerializeResponse(result, true)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	res := sr.(*types.JSONResponseSolo)

	return res, nil
}

// Delete
func (s *episodeService) Delete(ctx context.Context, id uuid.UUID) error {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Episode Service Delete called")

	if err := s.repo.Delete(ctx, id); err != nil {
		log.Error().Err(err)
		return err
	}

	return nil
}

// Detail
func (s *episodeService) Detail(ctx context.Context, id uuid.UUID) (*types.JSONResponseSolo, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Episode Service Detail called")

	result, err := s.repo.Detail(ctx, id)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	model := &types.Episode{}
	sr, err := model.SerializeResponse(result, true)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	res := sr.(*types.JSONResponseSolo)

	return res, nil
}

// List
func (s *episodeService) List(ctx context.Context, m *types.ListMeta) (*types.JSONResponseMult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Episode Service List called")

	return nil, nil // TODO
}

// Update
func (s *episodeService) Update(ctx context.Context, data any, id uuid.UUID) (*types.JSONResponseSolo, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := s.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Episode Service Update called")

	result, err := s.repo.Update(ctx, data.(*types.EpisodeRequestData), id)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	model := &types.Episode{}
	sr, err := model.SerializeResponse(result, true)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	res := sr.(*types.JSONResponseSolo)

	return res, nil
}
