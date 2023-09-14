package interfaces

import (
	"context"

	"github.com/BetterWorks/gosk-api/internal/core/models"
	"github.com/BetterWorks/gosk-api/internal/core/query"
	"github.com/google/uuid"
)

// ModelCreator
type ModelCreator interface {
	Create(context.Context, any) (DomainModel, error)
}

// ModelDeleter
type ModelDeleter interface {
	Delete(context.Context, uuid.UUID) error
}

// ModelDetailRetriever
type ModelDetailRetriever interface {
	Detail(context.Context, uuid.UUID) (DomainModel, error)
}

// ModelListRetriever
type ModelListRetriever interface {
	List(context.Context, query.QueryData) (DomainModel, error)
}

// ModelUpdater
type ModelUpdater interface {
	Update(context.Context, any, uuid.UUID) (DomainModel, error)
}

// ExampleRepository defines the interface for repository managing the Example domain/entity model
type ExampleRepository interface {
	Create(context.Context, *models.ExampleInputData) (DomainModel, error)
	ModelDeleter
	ModelDetailRetriever
	ModelListRetriever
	Update(context.Context, *models.ExampleInputData, uuid.UUID) (DomainModel, error)
}
