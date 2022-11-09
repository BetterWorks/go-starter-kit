package entities

import "github.com/jasonsites/gosk-api/internal/core/types"

type movieEntity struct{}

func NewMovieEntity() *movieEntity {
	return &movieEntity{}
}

func (r *movieEntity) Create(data *types.Movie) (*types.MovieRepoResult, error) {
	return &types.MovieRepoResult{}, nil
}

func (r *movieEntity) Delete(id string) error {
	return nil
}

func (r *movieEntity) Detail(id string) (*types.MovieRepoResult, error) {
	return &types.MovieRepoResult{}, nil
}

func (r *movieEntity) List(m *types.ListMeta) ([]*types.MovieRepoResult, error) {
	data := make([]*types.MovieRepoResult, 2)
	return data, nil
}

func (r *movieEntity) Update(data *types.Movie) (*types.MovieRepoResult, error) {
	return &types.MovieRepoResult{}, nil
}
