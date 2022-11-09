package entities

import (
	"fmt"

	"github.com/jasonsites/gosk-api/internal/core/types"
)

type bookEntity struct{}

func NewBookEntity() *bookEntity {
	return &bookEntity{}
}

func (r *bookEntity) Create(data *types.Book) (*types.BookRepoResult, error) {
	d := *data
	d.ID = "1234"
	entity := types.RepoEntity{Properties: d}

	result := &types.BookRepoResult{
		Data: []types.RepoEntity{entity},
	}
	fmt.Printf("Result in BookRepo.Create: %+v\n", result)

	return result, nil
}

func (r *bookEntity) Delete(id string) error {
	return nil
}

func (r *bookEntity) Detail(id string) (*types.BookRepoResult, error) {
	return &types.BookRepoResult{}, nil
}

func (r *bookEntity) List(m *types.ListMeta) ([]*types.BookRepoResult, error) {
	data := make([]*types.BookRepoResult, 2)
	return data, nil
}

func (r *bookEntity) Update(data *types.Book) (*types.BookRepoResult, error) {
	return &types.BookRepoResult{}, nil
}
