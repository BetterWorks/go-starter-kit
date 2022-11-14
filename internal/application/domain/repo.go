package domain

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
	Attributes any // TODO ?
	Related    []RepoResultRelatedResource
}

// RepoResultEntityMetadata
type RepoResultEntityMetadata struct{}

// RepoResultRelatedResource
type RepoResultRelatedResource struct {
	Type string
	Data []any // TODO
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

// EpisodeRepository
type EpisodeRepository interface {
	Create(*Episode) (*RepoResult, error)
	Delete(id string) error
	Detail(id string) (*RepoResult, error)
	List(*ListMeta) ([]*RepoResult, error)
	Update(*Episode) (*RepoResult, error)
}

// SeasonRepository
type SeasonRepository interface {
	Create(*Season) (*RepoResult, error)
	Delete(id string) error
	Detail(id string) (*RepoResult, error)
	List(*ListMeta) ([]*RepoResult, error)
	Update(*Season) (*RepoResult, error)
}
