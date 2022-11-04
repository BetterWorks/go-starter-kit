package application

import (
	"github.com/jasonsites/gosk-api/internal/core/types"
)

type Application struct {
	Repository types.Repository
}

func NewApplication(r types.Repository) *Application {
	return &Application{Repository: r}
}

// Create
func (a *Application) Create(data any) any {
	return data
}

// Delete
func (a *Application) Delete(id string) any {
	data := new(struct{})
	return data
}

// Detail
func (a *Application) Detail(id string) any {
	result := a.Repository.Detail(id)
	// s := serializers.NewSerializer()
	// return s.Serialize(result)
	return result
}

// List
func (a *Application) List(m *types.RequestMeta) any {
	data := new(struct{})
	return data
}

// Update
func (a *Application) Update(data any) any {
	return data
}
