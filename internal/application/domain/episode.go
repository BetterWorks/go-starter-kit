package domain

// Episode defines an example domain resource
type Episode struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        uint16 `json:"year"`
	Director    string `json:"director"`
	SeasonID    string `json:"season_id"`
	Status      int    `json:"status"`
	Deleted     bool   `json:"-"`
}

// Discover
func (m *Episode) Discover() *Episode {
	return m
}

// SerializeModel
func (m *Episode) SerializeModel(r *RepoResult, solo bool) (*Episode, error) {
	if solo {
		return r.Data[0].Attributes.(*Episode), nil
	}

	// TODO: List case
	return nil, nil
}

// SerializeResponse
func (m *Episode) SerializeResponse(r *RepoResult, solo bool) (JSONResponse, error) {
	if solo {
		model := r.Data[0].Attributes.(Episode)
		res := &JSONResponseSolo{
			Data: &ResponseResource{
				Type: DomainType.Episode,
				ID:   model.ID,
				Properties: &Episode{
					Description: model.Description,
					Director:    model.Director,
					SeasonID:    model.SeasonID,
					Status:      model.Status,
					Title:       model.Title,
					Year:        model.Year,
				},
			},
		}
		return res, nil
	}

	// TODO: List case
	res := &JSONResponseMult{
		Meta: &APIMetadata{
			Paging: &ListPaging{
				Limit:  r.Metadata.Paging.Limit,
				Offset: r.Metadata.Paging.Offset,
				Total:  r.Metadata.Paging.Total,
			},
		},
		Data: &[]ResponseResource{},
	}

	return res, nil
}
