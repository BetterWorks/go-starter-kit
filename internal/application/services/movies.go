package services

import (
	"fmt"

	"github.com/jasonsites/gosk-api/internal/core/types"
)

type movieService struct {
	Repo   types.MovieRepository
	logger *types.Logger
}

func NewMovieService(r types.MovieRepository) *movieService {
	return &movieService{
		Repo:   r,
		logger: nil,
	}
}

// Create
func (s *movieService) Create(data *types.Movie) (*types.JSONResponseDetail, error) {
	result, err := s.Repo.Create(data)
	if err != nil {
		fmt.Printf("Error in MovieService.Create on s.Repo.Create %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in MovieService.Create %+v\n", result)

	model := &types.Movie{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		fmt.Printf("Error in MovieService.Create on model.SerializeResponse %+v\n", err)
		return nil, err
	}
	r := res.(*types.JSONResponseDetail)
	fmt.Printf("Result in MovieService.Create on model.SerializeResponse (casted) %+v\n", r)

	return r, nil
}

// Delete
func (s *movieService) Delete(id string) error {
	if err := s.Repo.Delete(id); err != nil {
		fmt.Printf("Error in MovieService.Delete: %+v\n", err)
		return err
	}

	return nil
}

// Detail
func (s *movieService) Detail(id string) (*types.JSONResponseDetail, error) {
	result, err := s.Repo.Detail(id)
	if err != nil {
		fmt.Printf("Error in MovieService.Detail: %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in MovieService.Detail: %+v\n", result)

	model := &types.Movie{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		return nil, err
	}
	r := res.(*types.JSONResponseDetail)

	return r, nil
}

// List
func (s *movieService) List(m *types.ListMeta) (*types.JSONResponseList, error) {
	return nil, nil // TODO
}

// Update
func (s *movieService) Update(data *types.Movie) (*types.JSONResponseDetail, error) {
	result, err := s.Repo.Update(data)
	if err != nil {
		fmt.Printf("Error in MovieService.Update %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in MovieService.Update %+v\n", result)

	model := &types.Movie{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		return nil, err
	}
	r := res.(*types.JSONResponseDetail)

	return r, nil
}
