package types

// Resource defines the interface for all domain resources
type Resource interface {
	Discover() Resource
	// Serialize() (string, error)
}

// Resource Type Constants ------------------------------------------------------------------------
// ResourceRegistry defines a domain resource registry (constants map)
type resourceRegistry struct {
	Book  string
	Movie string
}

var ResourceType = resourceRegistry{
	Book:  "book",
	Movie: "movie",
}

// Resources -------------------------------------------------------------------------------------
// Book
type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  uint16 `json:"year"`
}

// NewBook
func NewBook() Resource {
	return new(Book)
}

// Discover
func (b *Book) Discover() Resource {
	return b
}

// -----------------------------------------------------------
// Movie defines an example domain resource
type Movie struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Year    uint16 `json:"year"`
	Enabled bool   `json:"enabled"`
	Deleted bool   `json:"deleted"`
	Status  int    `json:"status"`
}

// NewMovie
func NewMovie() Resource {
	return new(Movie)
}

// Discover
func (m *Movie) Discover() Resource {
	return m
}

// // Serialize
// func (r Movie) Serialize() (string, error) {
// 	s, err := json.Marshal(r)
// 	if err != nil {
// 		return "", fmt.Errorf("error marshalling json %v", err)
// 	}
// 	return string(s), nil
// }
