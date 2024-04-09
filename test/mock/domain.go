package mock

import (
	"context"
	"errors"

	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/core/pagination"
	"github.com/BetterWorks/go-starter-kit/internal/core/query"
	fx "github.com/BetterWorks/go-starter-kit/test/fixtures"
	"github.com/google/uuid"
)

// ExampleService
type ExampleService struct {
	CreateExampleError bool
	DeleteExampleError bool
	DetailExampleError bool
	ListExampleError   bool
	UpdateExampleError bool
	DetailResult       *models.ExampleDomainModel
	ListResult         *models.ExampleDomainModel
	CreateResult       *models.ExampleDomainModel
	UpdateResult       *models.ExampleDomainModel
}

func NewExampleService(config *ExampleService) *ExampleService {
	if config == nil {
		config = &ExampleService{}
	}

	if config.CreateResult == nil {
		config.CreateResult = &models.ExampleDomainModel{
			Data: []models.ExampleObject{fx.NewExampleObjectBuilder().Build()},
			Solo: true,
		}
	}

	if config.DetailResult == nil {
		config.DetailResult = &models.ExampleDomainModel{
			Data: []models.ExampleObject{fx.NewExampleObjectBuilder().Build()},
			Solo: true,
		}
	}

	if config.ListResult == nil {
		config.ListResult = &models.ExampleDomainModel{
			Data: []models.ExampleObject{},
			Meta: &models.ModelMetadata{
				Paging: pagination.PageMetadata{},
			},
			Solo: false,
		}
	}

	if config.UpdateResult == nil {
		config.UpdateResult = &models.ExampleDomainModel{
			Data: []models.ExampleObject{fx.NewExampleObjectBuilder().Build()},
			Solo: true,
		}
	}
	return &ExampleService{
		CreateExampleError: config.CreateExampleError,
		DeleteExampleError: config.DeleteExampleError,
		DetailExampleError: config.DetailExampleError,
		ListExampleError:   config.ListExampleError,
		UpdateExampleError: config.UpdateExampleError,
		DetailResult:       config.DetailResult,
		ListResult:         config.ListResult,
		CreateResult:       config.CreateResult,
		UpdateResult:       config.UpdateResult,
	}
}

// Create
func (s *ExampleService) Create(ctx context.Context, data *models.ExampleRequestAttributes) (*models.ExampleDomainModel, error) {
	if s.CreateExampleError {
		return nil, errors.New("service error")
	}

	return s.CreateResult, nil
}

// Delete
func (s *ExampleService) Delete(context.Context, uuid.UUID) error {
	if s.DeleteExampleError {
		return errors.New("service error")
	}

	return nil
}

// Detail
func (s *ExampleService) Detail(context.Context, uuid.UUID) (*models.ExampleDomainModel, error) {
	if s.DetailExampleError {
		return nil, errors.New("service error")
	}

	return s.DetailResult, nil
}

// List
func (s *ExampleService) List(context.Context, query.QueryData) (*models.ExampleDomainModel, error) {
	if s.ListExampleError {
		return nil, errors.New("service error")
	}

	return s.ListResult, nil
}

// Update
func (s *ExampleService) Update(context.Context, *models.ExampleRequestAttributes, uuid.UUID) (*models.ExampleDomainModel, error) {
	if s.UpdateExampleError {
		return nil, errors.New("service error")
	}

	return s.UpdateResult, nil
}
