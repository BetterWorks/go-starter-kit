package types

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// SeasonRequestData defines a Season domain model for data attributes request binding
type SeasonRequestData struct {
	Deleted     bool
	Description string
	Enabled     bool
	Status      uint16
	Title       string
}

// SeasonEntity defines a Season database entity
type SeasonEntity struct {
	CreatedBy   uint32
	CreatedOn   time.Time
	Deleted     bool
	Description sql.NullString
	Enabled     bool
	ID          uuid.UUID
	ModifiedBy  sql.NullInt32
	ModifiedOn  sql.NullTime
	Status      sql.NullInt16
	Title       string
}

// Season defines a Season domain model for application logic and response serialization
type Season struct {
	ID          uuid.UUID `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      uint16    `json:"status"`
	Enabled     bool      `json:"enabled"`
	Deleted     bool      `json:"-"`
	CreatedOn   time.Time `json:"created_on"`
	CreatedBy   uint32    `json:"created_by"`
	ModifiedOn  time.Time `json:"modified_on"`
	ModifiedBy  uint32    `json:"modified_by"`
}

// SerializeModel
func (s *Season) SerializeModel(r *RepoResult, solo bool) (*Season, error) {
	if solo {
		return r.Data[0].Attributes.(*Season), nil
	}

	// TODO: List case

	return nil, nil
}

// SerializeResponse
func (s *Season) SerializeResponse(r *RepoResult, solo bool) (JSONResponse, error) {
	if solo {
		model := r.Data[0].Attributes.(SeasonEntity)

		properties := &Season{
			CreatedBy:   model.CreatedBy,
			CreatedOn:   model.CreatedOn,
			Description: model.Description.String,
			Enabled:     model.Enabled,
			Status:      uint16(model.Status.Int16),
			Title:       model.Title,
		}

		result := &JSONResponseSolo{
			Data: &ResponseResource{
				Type:       DomainType.Season,
				ID:         model.ID,
				Properties: properties,
			},
		}

		return result, nil

	}

	meta := &APIMetadata{
		Paging: &ListPaging{
			Limit:  r.Metadata.Paging.Limit,
			Offset: r.Metadata.Paging.Offset,
			Total:  r.Metadata.Paging.Total,
		},
	}

	result := &JSONResponseMult{
		Meta: meta,
		Data: &[]ResponseResource{},
	}

	return result, nil
}
