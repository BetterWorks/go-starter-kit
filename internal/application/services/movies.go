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
	return nil, nil
}

// Delete
func (s *movieService) Delete(id string) error {
	if err := s.Repo.Delete(id); err != nil {
		fmt.Printf("Error in Movie Service Delete: %+v\n", err)
		return err
	}

	return nil
}

// Detail
func (s *movieService) Detail(id string) (*types.JSONResponseDetail, error) {
	result, err := s.Repo.Detail(id)
	if err != nil {
		fmt.Printf("Error in Movie Service Detail: %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in Movie Service Detail: %+v\n", result)

	// s := serializers.NewSerializer()
	// return s.Serialize(result)
	// return &types.Movie{}, nil
	return nil, nil
}

// List
func (s *movieService) List(m *types.ListMeta) (*types.JSONResponseList, error) {
	// data := make([]*types.Movie, 2)
	// return data, nil
	return nil, nil
}

// Update
func (s *movieService) Update(data *types.Movie) (*types.JSONResponseDetail, error) {
	return nil, nil
}
