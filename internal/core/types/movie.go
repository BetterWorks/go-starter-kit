package types

// Movie defines an example domain resource
type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Year     uint16 `json:"year"`
	Director string `json:"director"`
	Deleted  bool   `json:"deleted"`
	Status   int    `json:"status"`
}

func (m *Movie) Discover() *Movie {
	return m
}

// func (m *Movie) SerializeModel(r *MovieRepoResult) (*Movie, error) {
// 	return &Movie{}, nil
// }
