package types

// Book defines an example domain resource
type Book struct {
	ID string `json:"id"`
	BookProperties
}

// BookProperties
type BookProperties struct {
	Title   string `json:"title"`
	Year    uint16 `json:"year"`
	Author  string `json:"author"`
	Deleted bool   `json:"deleted"`
	Status  int    `json:"status"`
}

// Discover
func (b *Book) Discover() *Book {
	return b
}

// // SerializeModel
// func (b *Book) SerializeModel(r *BookRepoResult) (*Book, error) {
// 	return &Book{}, nil
// }

// SerializeResponse
func (b *Book) SerializeResponse(r *BookRepoResult, single bool) (JSONResponse, error) {
	if single {
		res := &JSONResponseDetail{
			Data: &Resource{
				Type:       "book",
				ID:         "1234",
				Properties: r.Data[0].Properties,
			},
		}
		return res, nil
	} else {
		res := &JSONResponseList{
			Meta: &APIMetadata{
				Paging: &ListPaging{
					Limit:  r.Metadata.Paging.Limit,
					Offset: r.Metadata.Paging.Offset,
					Total:  r.Metadata.Paging.Total,
				},
			},
			Data: &[]Resource{},
		}
		return res, nil
	}
}
