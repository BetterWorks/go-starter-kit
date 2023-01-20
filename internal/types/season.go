package types

import (
	"github.com/google/uuid"
)

// SeasonRequestData defines a Season domain model for data attributes request binding
type SeasonRequestData struct {
	Deleted     bool
	Description string
	Enabled     bool
	Status      uint8
	Title       string
}

// SeasonEntity defines a Season database entity
type SeasonEntity struct {
	Deleted     bool
	Description string
	Enabled     bool
	ID          string
	Status      int
	Title       string
}

// Season defines a Season domain model for application logic and response serialization
type Season struct {
	CreatedBy   uint32    `json:"created_by"`
	Deleted     bool      `json:"-"`
	Description string    `json:"description"`
	Enabled     bool      `json:"enabled"`
	ID          uuid.UUID `json:"id,omitempty"`
	Status      uint8     `json:"status"`
	Title       string    `json:"title"`
}

// Discover
func (s *Season) Discover() *Season {
	return s
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
		model := r.Data[0].Attributes.(Season)
		res := &JSONResponseSingle{
			Data: &ResponseResource{
				Type: DomainType.Season,
				ID:   model.ID,
				Properties: &Season{
					Description: model.Description,
					Title:       model.Title,
					Status:      model.Status,
				},
			},
		}
		return res, nil
	} else {
		res := &JSONResponseMulti{
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
}

// func (b *Season) Set(data Settable) Settable {
// 	s := data.(*Season)

// 	b.Deleted = s.Deleted
// 	b.Description = s.Description
// 	b.ID = s.ID
// 	b.Status = s.Status
// 	b.Title = s.Title

// 	return b
// }
