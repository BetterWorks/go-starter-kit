package types

// Discoverable
type Discoverable interface {
	Discover() Discoverable
}

// Model defines the interface for all domain resources
type Model interface {
	Discoverable
	// SerializeModel(any) (Model, error)
	SerializeResponse(any) (JSONResponse, error)
}
