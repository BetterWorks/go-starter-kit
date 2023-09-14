package interfaces

import "github.com/BetterWorks/gosk-api/internal/core/jsonapi"

// DomainModel defines the interface for all domain models
type DomainModel interface {
	ResponseFormatter
}

type ResponseFormatter interface {
	FormatResponse() (*jsonapi.Response, error)
}
