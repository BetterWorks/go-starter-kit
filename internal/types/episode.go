package types

import (
	"github.com/google/uuid"
)

// EpisodeRequestData defines an Episode domain model for data attributes request binding
type EpisodeRequestData struct {
	Deleted     bool
	Description string
	Director    string
	Enabled     bool
	SeasonID    uuid.UUID
	Status      uint8
	Title       string
	Year        uint16
}

// EpisodeEntity defines an Episode database entity
type EpisodeEntity struct {
	CreatedBy   uint32
	Deleted     bool
	Description string
	Director    string
	Enabled     bool
	ID          uuid.UUID
	SeasonID    uuid.UUID
	Status      uint8
	Title       string
	Year        uint16
}

// Episode defines an Episode domain model for application logic and response serialization
type Episode struct {
	ID          uuid.UUID `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Director    string    `json:"director"`
	Year        uint16    `json:"year"`
	SeasonID    uuid.UUID `json:"season_id,omitempty"`
	Status      uint8     `json:"status"`
	Enabled     bool      `json:"enabled"`
	Deleted     bool      `json:"-"`
	CreatedBy   uint32    `json:"created_by"`
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
		model := r.Data[0].Attributes.(EpisodeEntity)
		res := &JSONResponseSolo{
			Data: &ResponseResource{
				Type: DomainType.Episode,
				ID:   model.ID,
				Properties: &Episode{
					CreatedBy:   model.CreatedBy,
					Description: model.Description,
					Director:    model.Director,
					Enabled:     model.Enabled,
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
