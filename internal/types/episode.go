package types

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// EpisodeRequestData defines an Episode domain model for data attributes request binding
type EpisodeRequestData struct {
	Deleted     bool      `json:"deleted" validate:"omitempty,boolean"`
	Description *string   `json:"description" validate:"omitempty,min=3,max=999"`
	Director    *string   `json:"director" validate:"omitempty,min=2,max=255"`
	Enabled     bool      `json:"enabled" validate:"omitempty,boolean"`
	SeasonID    uuid.UUID `json:"season_id" validate:"required,uuid4"`
	Status      *uint32   `json:"status" validate:"omitempty,numeric"`
	Title       string    `json:"title" validate:"required,min=2,max=255"`
	Year        *uint32   `json:"year" validate:"omitempty,numeric"`
}

// EpisodeEntity defines an Episode database entity
type EpisodeEntity struct {
	CreatedBy   uint32
	CreatedOn   time.Time
	Deleted     bool
	Description sql.NullString
	Director    sql.NullString
	Enabled     bool
	ID          uuid.UUID
	ModifiedBy  sql.NullInt32
	ModifiedOn  sql.NullTime
	SeasonID    uuid.UUID
	Status      sql.NullInt32
	Title       string
	Year        sql.NullInt32
}

// Episode defines an Episode domain model for application logic and response serialization
type Episode struct {
	ID          uuid.UUID  `json:"-"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Director    *string    `json:"director"`
	Year        *uint32    `json:"year"`
	SeasonID    uuid.UUID  `json:"season_id"`
	Status      *uint32    `json:"status"`
	Enabled     bool       `json:"enabled"`
	Deleted     bool       `json:"-"`
	CreatedOn   time.Time  `json:"created_on"`
	CreatedBy   uint32     `json:"created_by"`
	ModifiedOn  *time.Time `json:"modified_on"`
	ModifiedBy  *uint32    `json:"modified_by"`
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
func (s *Episode) SerializeResponse(r *RepoResult, solo bool) (JSONResponse, error) {
	// single resource response
	if solo {
		resource := mapEpisodeEntityToResource(r.Data[0])
		result := &JSONResponseSolo{Data: resource}

		return result, nil
	}

	// multiple resource response
	meta := &APIMetadata{
		Paging: ListPaging{
			Limit:  r.Metadata.Paging.Limit,
			Offset: r.Metadata.Paging.Offset,
			Total:  r.Metadata.Paging.Total,
		},
	}

	data := make([]ResponseResource, 0)
	// TODO: go routine
	for _, record := range r.Data {
		resource := mapEpisodeEntityToResource(record)
		data = append(data, resource)
	}

	result := &JSONResponseMult{
		Meta: meta,
		Data: data,
	}

	return result, nil
}

// mapEpisodeEntityToResource maps an episode entity repo record to an episode response resource
func mapEpisodeEntityToResource(record RepoResultEntity) ResponseResource {
	model := record.Attributes.(EpisodeEntity)

	var (
		description *string
		director    *string
		modifiedBy  *uint32
		modifiedOn  *time.Time
		status      *uint32
		year        *uint32
	)

	if model.Description.Valid {
		description = &model.Description.String
	}
	if model.Director.Valid {
		director = &model.Director.String
	}
	if model.ModifiedBy.Valid {
		m := uint32(model.ModifiedBy.Int32)
		modifiedBy = &m
	}
	if model.ModifiedOn.Valid {
		modifiedOn = &model.ModifiedOn.Time
	}
	if model.Status.Valid {
		s := uint32(model.Status.Int32)
		status = &s
	}
	if model.Year.Valid {
		y := uint32(model.Year.Int32)
		year = &y
	}

	properties := &Episode{
		CreatedBy:   model.CreatedBy,
		CreatedOn:   model.CreatedOn,
		Description: description,
		Director:    director,
		Enabled:     model.Enabled,
		ModifiedBy:  modifiedBy,
		ModifiedOn:  modifiedOn,
		SeasonID:    model.SeasonID,
		Status:      status,
		Title:       model.Title,
		Year:        year,
	}

	return ResponseResource{
		Type:       DomainType.Episode,
		ID:         model.ID,
		Properties: properties,
	}
}
