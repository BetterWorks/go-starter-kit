package domain

import (
	"github.com/google/uuid"
)

// Season defines an example domain resource
type Season struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	Deleted     bool      `json:"-"`
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
					Title:       model.Title,
					Description: model.Description,
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
