package types

import (
	"github.com/google/uuid"
)

// Repository
type Repository interface {
	RepoCreator
	RepoDeleter
	RepoDetailRetriever
	RepoListRetriever
	RepoUpdater
}

// RepoCreator
type RepoCreator interface {
	Create(any) (*RepoResult, error)
}

// RepoDeleter
type RepoDeleter interface {
	Delete(id uuid.UUID) error
}

// RepoDetailRetriever
type RepoDetailRetriever interface {
	Detail(id uuid.UUID) (*RepoResult, error)
}

// RepoListRetriever
type RepoListRetriever interface {
	List(*ListMeta) ([]*RepoResult, error)
}

// RepoUpdater
type RepoUpdater interface {
	Update(any) (*RepoResult, error)
}

// temp documentation
// {
// 		meta: {
// 				paging: {
// 						limit,
// 						offset,
// 						total
// 				}
// 		},
// 		data: [{
// 				type: 'resource-type',
// 				meta: {
// 						...resource metadata
// 				},
// 				properties: {
// 						...resource properties
// 				},
// 				rel: [{
// 						type: 'rel-type',
// 						data: [{
// 								...rel-resource
// 						}],
// 				}],
// 		}]
// }

// RepoResult
type RepoResult struct {
	Metadata RepoResultMetadata
	Data     []RepoResultEntity
}

// RepoResultMetadata
type RepoResultMetadata struct {
	Paging ListPaging
}

// RepoResultEntity
type RepoResultEntity struct {
	Type       string
	Meta       RepoResultEntityMetadata
	Attributes any
	Related    []RepoResultRelatedEntity
}

// RepoResultEntityMetadata
type RepoResultEntityMetadata struct{}

// RepoResultRelatedEntity
type RepoResultRelatedEntity struct {
	Type string
	Data []any // TODO
}
