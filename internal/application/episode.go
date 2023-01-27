package application

import (
	"context"
	"fmt"

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
	result, err := s.repo.Create(ctx, data.(*types.EpisodeRequestData))
	if err != nil {
		fmt.Printf("Error in episodeService.Create on s.repo.Create %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in episodeService.Create %+v\n", result)

	model := &types.Episode{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		fmt.Printf("Error in episodeService.Create on model.SerializeResponse %+v\n", err)
		return nil, err
	}
	r := res.(*types.JSONResponseSolo)
	fmt.Printf("Result in episodeService.Create on model.SerializeResponse (casted) %+v\n", r)

	return r, nil
}

// Delete
func (s *episodeService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		fmt.Printf("Error in episodeService.Delete: %+v\n", err)
		return err
	}

	return nil
}

// Detail
func (s *episodeService) Detail(ctx context.Context, id uuid.UUID) (*types.JSONResponseSolo, error) {
	result, err := s.repo.Detail(ctx, id)
	if err != nil {
		fmt.Printf("Error in episodeService.Detail: %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in episodeService.Detail: %+v\n", result)

	model := &types.Episode{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		return nil, err
	}
	r := res.(*types.JSONResponseSolo)

	return r, nil
}

// List
func (s *episodeService) List(ctx context.Context, m *types.ListMeta) (*types.JSONResponseMult, error) {
	return nil, nil // TODO
}

// Update
func (s *episodeService) Update(ctx context.Context, data any, id uuid.UUID) (*types.JSONResponseSolo, error) {
	result, err := s.repo.Update(ctx, data.(*types.EpisodeRequestData), id)
	if err != nil {
		fmt.Printf("Error in episodeService.Update %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in episodeService.Update %+v\n", result)

	model := &types.Episode{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		return nil, err
	}
	r := res.(*types.JSONResponseSolo)

	return r, nil
}
