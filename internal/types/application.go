package types

import (
	"context"

	"github.com/google/uuid"
)

// DomainModel defines the interface for all domain models
type DomainModel interface {
	// Discoverable
	ResponseSerializer
}

// domainRegistry defines a registry for all domain types to be used across the application
type DomainRegistry struct {
	Episode string
	Season  string
}

// DomainType exposes constants for all domain types
var DomainType = DomainRegistry{
	Episode: "episode",
	Season:  "season",
}

// TODO
// Discoverable defines the interface for all types with self discovery
type Discoverable interface {
	Discover() Discoverable
}

// TODO
// ResponseSerializer defines the interface for all types that serialize to JSON response
type ResponseSerializer interface {
	SerializeResponse(any, bool) (JSONResponse, error)
}

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
	List(context.Context, QueryData) (*JSONResponseMult, error)
}

type ServiceUpdater interface {
	Update(context.Context, any, uuid.UUID) (*JSONResponseSolo, error)
}
