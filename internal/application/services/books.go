package services

import (
	"fmt"

	"github.com/jasonsites/gosk-api/internal/core/types"
)

type bookService struct {
	Repo   types.BookRepository
	logger *types.Logger
}

func NewBookService(r types.BookRepository) *bookService {
	return &bookService{
		Repo:   r,
		logger: nil,
	}
}

// Create
func (s *bookService) Create(data *types.Book) (*types.JSONResponseDetail, error) {
	result, err := s.Repo.Create(data)
	if err != nil {
		fmt.Printf("Error in BookService.Create on s.Repo.Create %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in BookService.Create %+v\n", result)

	model := &types.Book{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		fmt.Printf("Error in BookService.Create on model.SerializeResponse %+v\n", err)
		return nil, err
	}
	r := res.(*types.JSONResponseDetail)
	fmt.Printf("Result in BookService.Create on model.SerializeResponse (casted) %+v\n", r)

	return r, nil
}

// Delete
func (s *bookService) Delete(id string) error {
	if err := s.Repo.Delete(id); err != nil {
		fmt.Printf("Error in BookService.Delete: %+v\n", err)
		return err
	}

	return nil
}

// Detail
func (s *bookService) Detail(id string) (*types.JSONResponseDetail, error) {
	result, err := s.Repo.Detail(id)
	if err != nil {
		fmt.Printf("Error in BookService.Detail: %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in BookService.Detail: %+v\n", result)

	model := &types.Book{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		return nil, err
	}
	r := res.(*types.JSONResponseDetail)

	return r, nil
}

// List
func (s *bookService) List(m *types.ListMeta) (*types.JSONResponseList, error) {
	return nil, nil
}

// Update
func (s *bookService) Update(data *types.Book) (*types.JSONResponseDetail, error) {
	result, err := s.Repo.Update(data)
	if err != nil {
		fmt.Printf("Error in BookService.Update %+v\n", err)
		return nil, err
	}
	fmt.Printf("Result in BookService.Update %+v\n", result)

	model := &types.Book{}
	res, err := model.SerializeResponse(result, true)
	if err != nil {
		// log error
		return nil, err
	}
	r := res.(*types.JSONResponseDetail)

	return r, nil
}
