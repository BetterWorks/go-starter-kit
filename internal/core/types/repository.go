package types

type RepoResultMetadata struct {
	Paging ListPaging // `json:"paging,omitempty"`
}

type RepoResultData struct {
	Data []RepoEntity // `json:"data"`
}

type RepoEntity struct {
	Type string // `json:"type"`
	// Meta       ResourceMetadata      `json:"meta"`
	Properties any // `json:"properties"`
	// Related    []RepoRelatedResource `json:"rel"`
}

type RepoRelatedResource struct {
	Type string `json:"type"`
	Data []any  `json:"data"` // TODO
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

// BookRepoResult
type BookRepoResult struct {
	Metadata RepoResultMetadata // `json:"meta"`
	Data     []RepoEntity       // `json:"data"`
}

// BookRepository
type BookRepository interface {
	Create(*Book) (*BookRepoResult, error)
	Delete(id string) error
	Detail(id string) (*BookRepoResult, error)
	List(*ListMeta) ([]*BookRepoResult, error)
	Update(*Book) (*BookRepoResult, error)
}

// MovieRepoResult
type MovieRepoResult struct{}

// MovieRepository
type MovieRepository interface {
	Create(*Movie) (*MovieRepoResult, error)
	Delete(id string) error
	Detail(id string) (*MovieRepoResult, error)
	List(*ListMeta) ([]*MovieRepoResult, error)
	Update(*Movie) (*MovieRepoResult, error)
}
