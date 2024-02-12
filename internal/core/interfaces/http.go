package interfaces

import (
	"net/http"
)

// ExampleController
type ExampleController interface {
	Create() http.HandlerFunc
	Delete() http.HandlerFunc
	Detail() http.HandlerFunc
	List() http.HandlerFunc
	Update() http.HandlerFunc
}
