package domain

import (
	"github.com/google/uuid"
)

// Repository
type Repository interface {
	Create(any) (*RepoResult, error)
	Delete(id uuid.UUID) error
	Detail(id uuid.UUID) (*RepoResult, error)
	List(*ListMeta) ([]*RepoResult, error)
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
