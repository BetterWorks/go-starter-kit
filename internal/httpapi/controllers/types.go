package controllers

// JSONRequestBody
type JSONRequestBody struct {
	Data *RequestResource
}

// RequestResource
type RequestResource struct {
	Type       string
	ID         string
	Properties any
}
