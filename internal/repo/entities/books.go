package entities

import (
	"fmt"

	"github.com/jasonsites/gosk-api/internal/types"
)

// bookEntity
type bookEntity struct {
	Author  string
	ID      string
	Deleted bool
	Status  int
	Title   string
	Year    uint16
}

// NewBookEntity
func NewBookEntity() *bookEntity {
	return &bookEntity{}
}

// Create
func (r *bookEntity) Create(data *types.Book) (*types.BookRepoResult, error) {
	d := *data
	d.ID = "1234"
	entity := types.RepoResultEntity{Properties: d}

	result := &types.BookRepoResult{
		Data: []types.RepoResultEntity{entity},
	}
	fmt.Printf("Result in BookRepo.Create: %+v\n", result)

	return result, nil
}

// Delete
func (r *bookEntity) Delete(id string) error {
	return nil
}

// Detail
func (r *bookEntity) Detail(id string) (*types.BookRepoResult, error) {
	return &types.BookRepoResult{}, nil
}

// List
func (r *bookEntity) List(m *types.ListMeta) ([]*types.BookRepoResult, error) {
	data := make([]*types.BookRepoResult, 2)
	return data, nil
}

// Update
func (r *bookEntity) Update(data *types.Book) (*types.BookRepoResult, error) {
	return &types.BookRepoResult{}, nil
}
