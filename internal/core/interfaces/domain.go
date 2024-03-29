package interfaces

import (
	"context"

	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/core/query"
	"github.com/google/uuid"
)

// ExampleService
type ExampleService interface {
	Create(context.Context, *models.ExampleRequestAttributes) (*models.ExampleDomainModel, error)
	Delete(context.Context, uuid.UUID) error
	Detail(context.Context, uuid.UUID) (*models.ExampleDomainModel, error)
	List(context.Context, query.QueryData) (*models.ExampleDomainModel, error)
	Update(context.Context, *models.ExampleRequestAttributes, uuid.UUID) (*models.ExampleDomainModel, error)
}
