package types

// Discoverable
type Discoverable interface {
	Discover() Discoverable
}

type ResponseSerializer interface {
	SerializeResponse(any, bool) (JSONResponse, error)
}

type Settable interface {
	Set(Settable) Settable
}

// Model defines the interface for all domain models
type Model interface {
	Discoverable
	ResponseSerializer
	Settable
}

// domainRegistry defines a domain registry (constants)
type domainRegistry struct {
	Book  string
	Movie string
}

// DomainType exposes constants for all domain types
var DomainType = domainRegistry{
	Book:  "book",
	Movie: "movie",
}
