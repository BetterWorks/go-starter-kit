package types

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// EpisodeRequestData defines an Episode domain model for data attributes request binding
type EpisodeRequestData struct {
	Deleted     bool
	Description string
	Director    string
	Enabled     bool
	SeasonID    uuid.UUID
	Status      uint32
	Title       string
	Year        uint32
}

// EpisodeEntity defines an Episode database entity
type EpisodeEntity struct {
	CreatedBy   uint32
	CreatedOn   time.Time
	Deleted     bool
	Description string
	Director    string
	Enabled     bool
	ID          uuid.UUID
	ModifiedBy  sql.NullInt32
	ModifiedOn  sql.NullTime
	SeasonID    uuid.UUID
	Status      sql.NullInt32
	Title       string
	Year        uint32
}

// Episode defines an Episode domain model for application logic and response serialization
type Episode struct {
	ID          uuid.UUID `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Director    string    `json:"director"`
	Year        uint32    `json:"year"`
	SeasonID    uuid.UUID `json:"season_id,omitempty"`
	Status      uint32    `json:"status"`
	Enabled     bool      `json:"enabled"`
	Deleted     bool      `json:"-"`
	CreatedOn   time.Time `json:"created_on"`
	CreatedBy   uint32    `json:"created_by"`
	ModifiedOn  time.Time `json:"modified_on"`
	ModifiedBy  uint32    `json:"modified_by"`
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
					Status:      uint32(model.Status.Int32),
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
