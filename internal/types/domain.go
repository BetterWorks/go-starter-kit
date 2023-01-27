package types

// Discoverable defines the interface for all types with self discovery
type Discoverable interface {
	Discover() Discoverable
}

// ResponseSerializer defines the interface for all types that serialize to JSON response
type ResponseSerializer interface {
	SerializeResponse(any, bool) (JSONResponse, error)
}

// DomainModel defines the interface for all domain models
type DomainModel interface {
	// Discoverable
	ResponseSerializer
}

// domainRegistry defines a domain registry (constants)
type DomainRegistry struct {
	Episode string
	Season  string
}

// DomainType exposes constants for all domain types
var DomainType = DomainRegistry{
	Episode: "episode",
	Season:  "season",
}
