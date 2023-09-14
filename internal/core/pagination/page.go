package pagination

// PageMetadata
type PageMetadata struct {
	Limit  uint32 `json:"limit"`
	Offset uint32 `json:"offset"`
	Total  uint32 `json:"total"`
}

// SortMetadata
type SortMetadata struct {
	Attr  string
	Order string // "asc" || "desc"
}
