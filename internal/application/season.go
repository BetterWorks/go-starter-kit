package application

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jasonsites/gosk-api/internal/types"
)

type seasonService struct {
	Repo   types.Repository
	logger *types.Logger
}

func NewSeasonService(r types.Repository) *seasonService {
	return &seasonService{
		Repo:   r,
		logger: nil,
	}
}

// Create
func (s *seasonService) Create(ctx context.Context, data any) (*types.JSONResponseSolo, error) {
	result, err := s.Repo.Create(ctx, data.(*types.SeasonRequestData))
	if err != nil {
		fmt.Printf("Error in seasonService.Create on s.Repo.Create %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in seasonService.Create %+v\n", result)

	model := &types.Season{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		fmt.Printf("Error in seasonService.Create on model.SerializeResponse %+v\n", err)
		return nil, err
	}
	r := res.(*types.JSONResponseSolo)
	fmt.Printf("Result in seasonService.Create on model.SerializeResponse (casted) %+v\n", r)

	return r, nil
}

// Delete
func (s *seasonService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.Repo.Delete(ctx, id); err != nil {
		fmt.Printf("Error in seasonService.Delete: %+v\n", err)
		return err
	}

	return nil
}

// Detail
func (s *seasonService) Detail(ctx context.Context, id uuid.UUID) (*types.JSONResponseSolo, error) {
	result, err := s.Repo.Detail(ctx, id)
	if err != nil {
		fmt.Printf("Error in seasonService.Detail: %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in seasonService.Detail: %+v\n", result)

	model := &types.Season{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		fmt.Printf("Error in seasonService.Detail: %+v\n", err)
		return nil, err
	}
	r := res.(*types.JSONResponseSolo)

	return r, nil
}

// List
func (s *seasonService) List(ctx context.Context, m *types.ListMeta) (*types.JSONResponseMult, error) {
	return nil, nil // TODO
}

// Update
func (s *seasonService) Update(ctx context.Context, data any) (*types.JSONResponseSolo, error) {
	result, err := s.Repo.Update(ctx, data.(*types.SeasonRequestData))
	if err != nil {
		fmt.Printf("Error in seasonService.Update %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in seasonService.Update %+v\n", result)

	model := &types.Season{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		fmt.Printf("Error in seasonService.Update on model.SerializeResponse %+v\n", err)
		return nil, err
	}
	r := res.(*types.JSONResponseSolo)

	return r, nil
}
