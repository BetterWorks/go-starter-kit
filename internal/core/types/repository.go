package types

// REPOSITORY -------------------------------------------------------------------------------------
// RepoResult
type RepoResult struct{}

// Repository
type Repository interface {
	Create(data any) *RepoResult
	Delete(id string) *RepoResult
	Detail(id string) *RepoResult
	List(*RequestMeta) *RepoResult
	Update(data any) *RepoResult
}
