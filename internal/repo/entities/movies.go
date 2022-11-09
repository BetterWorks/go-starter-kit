package entities

import (
	"fmt"

	"github.com/jasonsites/gosk-api/internal/core/types"
)

// movieEntity
type movieEntity struct {
	Deleted  bool
	Director string
	ID       string
	Status   int
	Title    string
	Year     uint16
}

// NewMovieEntity
func NewMovieEntity() *movieEntity {
	return &movieEntity{}
}

// Create
func (r *movieEntity) Create(data *types.Movie) (*types.MovieRepoResult, error) {
	d := *data
	d.ID = "1234"
	entity := types.RepoResultEntity{Properties: d}

	result := &types.MovieRepoResult{
		Data: []types.RepoResultEntity{entity},
	}
	fmt.Printf("Result in MovieRepo.Create: %+v\n", result)

	return result, nil
}

// Delete
func (r *movieEntity) Delete(id string) error {
	return nil
}

// Detail
func (r *movieEntity) Detail(id string) (*types.MovieRepoResult, error) {
	return &types.MovieRepoResult{}, nil
}

// List
func (r *movieEntity) List(m *types.ListMeta) ([]*types.MovieRepoResult, error) {
	data := make([]*types.MovieRepoResult, 2)
	return data, nil
}

// Update
func (r *movieEntity) Update(data *types.Movie) (*types.MovieRepoResult, error) {
	return &types.MovieRepoResult{}, nil
}
