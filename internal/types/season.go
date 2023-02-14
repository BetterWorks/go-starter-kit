package types

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// SeasonRequestData defines a Season domain model for data attributes request binding
type SeasonRequestData struct {
	Deleted     bool    `json:"deleted"`
	Description *string `json:"description"`
	Enabled     bool    `json:"enabled"`
	Status      *uint32 `json:"status"`
	Title       string  `json:"title"`
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
	Status      sql.NullInt32
	Title       string
}

// Season defines a Season domain model for application logic and response serialization
type Season struct {
	ID          uuid.UUID  `json:"-"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      *uint32    `json:"status"`
	Enabled     bool       `json:"enabled"`
	Deleted     bool       `json:"-"`
	CreatedOn   time.Time  `json:"created_on"`
	CreatedBy   uint32     `json:"created_by"`
	ModifiedOn  *time.Time `json:"modified_on"`
	ModifiedBy  *uint32    `json:"modified_by"`
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
	// single resource response
	if solo {
		resource := mapSeasonEntityToResource(r.Data[0])
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
		resource := mapSeasonEntityToResource(record)
		data = append(data, resource)
	}

	result := &JSONResponseMult{
		Meta: meta,
		Data: data,
	}

	return result, nil
}

// mapSeasonEntityToResource maps a season entity repo record to a season response resource
func mapSeasonEntityToResource(record RepoResultEntity) ResponseResource {
	model := record.Attributes.(SeasonEntity)

	var (
		description *string
		modifiedBy  *uint32
		modifiedOn  *time.Time
		status      *uint32
	)

	if model.Description.Valid {
		description = &model.Description.String
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

	properties := &Season{
		CreatedBy:   model.CreatedBy,
		CreatedOn:   model.CreatedOn,
		Description: description,
		Enabled:     model.Enabled,
		ModifiedBy:  modifiedBy,
		ModifiedOn:  modifiedOn,
		Status:      status,
		Title:       model.Title,
	}

	return ResponseResource{
		Type:       DomainType.Season,
		ID:         model.ID,
		Properties: properties,
	}
}
