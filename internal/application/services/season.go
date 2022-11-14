package services

import (
	"fmt"

	"github.com/jasonsites/gosk-api/internal/application/domain"
)

type seasonService struct {
	Repo   domain.SeasonRepository
	logger *domain.Logger
}

func NewSeasonService(r domain.SeasonRepository) *seasonService {
	return &seasonService{
		Repo:   r,
		logger: nil,
	}
}

// Create
func (s *seasonService) Create(data any) (*domain.JSONResponseSolo, error) {
	result, err := s.Repo.Create(data.(*domain.Season))
	if err != nil {
		fmt.Printf("Error in seasonService.Create on s.Repo.Create %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in seasonService.Create %+v\n", result)

	model := &domain.Season{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		fmt.Printf("Error in seasonService.Create on model.SerializeResponse %+v\n", err)
		return nil, err
	}
	r := res.(*domain.JSONResponseSolo)
	fmt.Printf("Result in seasonService.Create on model.SerializeResponse (casted) %+v\n", r)

	return r, nil
}

// Delete
func (s *seasonService) Delete(id string) error {
	if err := s.Repo.Delete(id); err != nil {
		fmt.Printf("Error in seasonService.Delete: %+v\n", err)
		return err
	}

	return nil
}

// Detail
func (s *seasonService) Detail(id string) (*domain.JSONResponseSolo, error) {
	result, err := s.Repo.Detail(id)
	if err != nil {
		fmt.Printf("Error in seasonService.Detail: %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in seasonService.Detail: %+v\n", result)

	model := &domain.Season{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		fmt.Printf("Error in seasonService.Detail: %+v\n", err)
		return nil, err
	}
	r := res.(*domain.JSONResponseSolo)

	return r, nil
}

// List
func (s *seasonService) List(m *domain.ListMeta) (*domain.JSONResponseMult, error) {
	return nil, nil // TODO
}

// Update
func (s *seasonService) Update(data any) (*domain.JSONResponseSolo, error) {
	result, err := s.Repo.Update(data.(*domain.Season))
	if err != nil {
		fmt.Printf("Error in seasonService.Update %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in seasonService.Update %+v\n", result)

	model := &domain.Season{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		fmt.Printf("Error in seasonService.Update on model.SerializeResponse %+v\n", err)
		return nil, err
	}
	r := res.(*domain.JSONResponseSolo)

	return r, nil
}
