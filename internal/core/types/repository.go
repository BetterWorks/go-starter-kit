package types

// RepoResultMetadata
type RepoResultMetadata struct {
	Paging ListPaging // `json:"paging,omitempty"`
}

// RepoResultEntity
type RepoResultEntity struct {
	Type       string
	Meta       RepoResultEntityMetadata
	Properties any // TODO
	Related    []RepoResultRelatedResource
}

// RepoResultEntityMetadata
type RepoResultEntityMetadata struct{}

// RepoResultRelatedResource
type RepoResultRelatedResource struct {
	Type string
	Data []any // TODO
}

// BookRepoResult
type BookRepoResult struct {
	Metadata RepoResultMetadata
	Data     []RepoResultEntity
}

// MovieRepoResult
type MovieRepoResult struct {
	Metadata RepoResultMetadata
	Data     []RepoResultEntity
}

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

// BookRepository
type BookRepository interface {
	Create(*Book) (*BookRepoResult, error)
	Delete(id string) error
	Detail(id string) (*BookRepoResult, error)
	List(*ListMeta) ([]*BookRepoResult, error)
	Update(*Book) (*BookRepoResult, error)
}

// MovieRepository
type MovieRepository interface {
	Create(*Movie) (*MovieRepoResult, error)
	Delete(id string) error
	Detail(id string) (*MovieRepoResult, error)
	List(*ListMeta) ([]*MovieRepoResult, error)
	Update(*Movie) (*MovieRepoResult, error)
}
