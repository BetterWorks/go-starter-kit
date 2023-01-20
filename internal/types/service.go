package types

import "github.com/google/uuid"

// Service
type Service interface {
	ServiceCreator
	ServiceDeleter
	ServiceDetailRetriever
	ServiceListRetriever
	ServiceUpdater
}

type ServiceCreator interface {
	Create(any) (*JSONResponseSingle, error)
}

type ServiceDeleter interface {
	Delete(uuid.UUID) error
}

type ServiceDetailRetriever interface {
	Detail(uuid.UUID) (*JSONResponseSingle, error)
}

type ServiceListRetriever interface {
	List(*ListMeta) (*JSONResponseMulti, error)
}

type ServiceUpdater interface {
	Update(any) (*JSONResponseSingle, error)
}
