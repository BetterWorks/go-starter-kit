package types

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// ResourceRequestData defines a Resource domain model for data attributes request binding
type ResourceRequestData struct {
	Deleted     bool    `json:"deleted" validate:"omitempty,boolean"`
	Description *string `json:"description" validate:"omitempty,min=3,max=999"`
	Enabled     bool    `json:"enabled"  validate:"omitempty,boolean"`
	Status      *uint32 `json:"status" validate:"omitempty,numeric"`
	Title       string  `json:"title" validate:"required,omitempty,min=2,max=255"`
}

// ResourceEntity defines a Resource database entity
type ResourceEntity struct {
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

// Resource defines a Resource domain model for application logic and response serialization
type Resource struct {
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
func (s *Resource) SerializeModel(r *RepoResult, solo bool) (*Resource, error) {
	if solo {
		return r.Data[0].Attributes.(*Resource), nil
	}

	// TODO: List case

	return nil, nil
}

// SerializeResponse
func (s *Resource) SerializeResponse(r *RepoResult, solo bool) (JSONResponse, error) {
	// single resource response
	if solo {
		resource := mapResourceEntityToResource(r.Data[0])
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
		resource := mapResourceEntityToResource(record)
		data = append(data, resource)
	}

	result := &JSONResponseMult{
		Meta: meta,
		Data: data,
	}

	return result, nil
}

// mapResourceEntityToResource maps a resource entity repo record to a resource response resource
func mapResourceEntityToResource(record RepoResultEntity) ResponseResource {
	model := record.Attributes.(ResourceEntity)

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

	properties := &Resource{
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
		Type:       DomainType.Resource,
		ID:         model.ID,
		Properties: properties,
	}
}
