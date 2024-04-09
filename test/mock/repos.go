package mock

import (
	"context"
	"errors"

	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/core/query"
	"github.com/google/uuid"
)

type ExampleRepository struct {
	CreateExampleError bool
	DeleteExampleError bool
	DetailExampleError bool
	ListExampleError   bool
	UpdateExampleError bool
}

// Create
func (r *ExampleRepository) Create(context.Context, *models.ExampleRequestAttributes) (*models.ExampleDomainModel, error) {
	if r.CreateExampleError {
		return nil, errors.New("repository error")
	}

	return &models.ExampleDomainModel{}, nil
}

// Delete
func (r *ExampleRepository) Delete(context.Context, uuid.UUID) error {
	if r.DeleteExampleError {
		return errors.New("repository error")
	}

	return nil
}

// Detail
func (r *ExampleRepository) Detail(context.Context, uuid.UUID) (*models.ExampleDomainModel, error) {
	if r.DetailExampleError {
		return nil, errors.New("repository error")
	}

	return &models.ExampleDomainModel{}, nil
}

// List
func (r *ExampleRepository) List(context.Context, query.QueryData) (*models.ExampleDomainModel, error) {
	if r.ListExampleError {
		return nil, errors.New("repository error")
	}

	return &models.ExampleDomainModel{}, nil
}

// Update
func (r *ExampleRepository) Update(context.Context, *models.ExampleRequestAttributes, uuid.UUID) (*models.ExampleDomainModel, error) {
	if r.UpdateExampleError {
		return nil, errors.New("repository error")
	}

	return &models.ExampleDomainModel{}, nil
}
