package interfaces

import (
	"context"

	"github.com/BetterWorks/gosk-api/internal/core/query"
	"github.com/google/uuid"
)

// Service
type Service interface {
	ServiceCreator
	ServiceDeleter
	ServiceDetailRetriever
	ServiceListRetriever
	ServiceUpdater
}

type ServiceCreator interface {
	Create(context.Context, any) (DomainModel, error)
}

type ServiceDeleter interface {
	Delete(context.Context, uuid.UUID) error
}

type ServiceDetailRetriever interface {
	Detail(context.Context, uuid.UUID) (DomainModel, error)
}

type ServiceListRetriever interface {
	List(context.Context, query.QueryData) (DomainModel, error)
}

type ServiceUpdater interface {
	Update(context.Context, any, uuid.UUID) (DomainModel, error)
}
