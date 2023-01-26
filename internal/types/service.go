package types

import (
	"context"

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
	Create(context.Context, any) (*JSONResponseSolo, error)
}

type ServiceDeleter interface {
	Delete(context.Context, uuid.UUID) error
}

type ServiceDetailRetriever interface {
	Detail(context.Context, uuid.UUID) (*JSONResponseSolo, error)
}

type ServiceListRetriever interface {
	List(context.Context, *ListMeta) (*JSONResponseMult, error)
}

type ServiceUpdater interface {
	Update(context.Context, any) (*JSONResponseSolo, error)
}
