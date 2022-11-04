package repo

import "github.com/jasonsites/gosk-api/internal/core/types"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(data any) *types.RepoResult {
	return &types.RepoResult{}
}

func (r *Repository) Delete(id string) *types.RepoResult {
	return &types.RepoResult{}
}

func (r *Repository) Detail(id string) *types.RepoResult {
	// TODO: 1. design RepoResult 2. implement actual deep repo db access
	return &types.RepoResult{}
}

func (r *Repository) List(m *types.RequestMeta) *types.RepoResult {
	return &types.RepoResult{}
}

func (r *Repository) Update(data any) *types.RepoResult {
	return &types.RepoResult{}
}
