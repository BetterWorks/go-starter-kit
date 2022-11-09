package types

// Discoverable
type Discoverable interface {
	Discover() Discoverable
}

// Model defines the interface for all domain models
type Model interface {
	Discoverable
	SerializeResponse(any) (JSONResponse, error)
}
