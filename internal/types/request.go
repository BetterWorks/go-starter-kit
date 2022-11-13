package types

// JSON API Request ------------------------------------------------------------------------------
// // JSONRequest defines the interface for a JSON REST request body
// type JSONRequest interface {
// 	Discoverable
// 	Settable
// }

// RequestResource
type RequestResource struct {
	Type       string
	ID         string
	Properties any
}

// JSONRequestBody
type JSONRequestBody struct {
	Data *RequestResource
}

// // Discover
// func (r *JSONRequestBody) Discover() Discoverable {
// 	return r
// }

// type BookRequestBody struct {
// 	Data *RequestResource
// }

// func (b *BookRequestBody) Set(req *JSONRequestBody) *JSONRequestBody {
// 	req.Data.Properties = &Book{}
// 	return req
// }
