package services

import (
	"fmt"

	"github.com/jasonsites/gosk-api/internal/application/domain"
)

type episodeService struct {
	Repo   domain.EpisodeRepository
	logger *domain.Logger
}

func NewEpisodeService(r domain.EpisodeRepository) *episodeService {
	return &episodeService{
		Repo:   r,
		logger: nil,
	}
}

// Create
func (s *episodeService) Create(data any) (*domain.JSONResponseSolo, error) {
	result, err := s.Repo.Create(data.(*domain.Episode))
	if err != nil {
		fmt.Printf("Error in episodeService.Create on s.Repo.Create %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in episodeService.Create %+v\n", result)

	model := &domain.Episode{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		fmt.Printf("Error in episodeService.Create on model.SerializeResponse %+v\n", err)
		return nil, err
	}
	r := res.(*domain.JSONResponseSolo)
	fmt.Printf("Result in episodeService.Create on model.SerializeResponse (casted) %+v\n", r)

	return r, nil
}

// Delete
func (s *episodeService) Delete(id string) error {
	if err := s.Repo.Delete(id); err != nil {
		fmt.Printf("Error in episodeService.Delete: %+v\n", err)
		return err
	}

	return nil
}

// Detail
func (s *episodeService) Detail(id string) (*domain.JSONResponseSolo, error) {
	result, err := s.Repo.Detail(id)
	if err != nil {
		fmt.Printf("Error in episodeService.Detail: %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in episodeService.Detail: %+v\n", result)

	model := &domain.Episode{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		return nil, err
	}
	r := res.(*domain.JSONResponseSolo)

	return r, nil
}

// List
func (s *episodeService) List(m *domain.ListMeta) (*domain.JSONResponseMult, error) {
	return nil, nil // TODO
}

// Update
func (s *episodeService) Update(data any) (*domain.JSONResponseSolo, error) {
	result, err := s.Repo.Update(data.(*domain.Episode))
	if err != nil {
		fmt.Printf("Error in episodeService.Update %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in episodeService.Update %+v\n", result)

	model := &domain.Episode{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		return nil, err
	}
	r := res.(*domain.JSONResponseSolo)

	return r, nil
}
