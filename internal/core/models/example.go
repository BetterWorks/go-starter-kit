package models

import (
	"time"

	"github.com/BetterWorks/go-starter-kit/internal/core/pagination"
	"github.com/google/uuid"
	v "github.com/invopop/validation"
)

// ExampleRequest
type ExampleRequest struct {
	Data *ExampleRequestResource `json:"data" validate:"required"`
}

// ExampleRequestResource
type ExampleRequestResource struct {
	Type       string                   `json:"type" validate:"required"`
	ID         string                   `json:"id" validate:"omitempty,uuid4"`
	Attributes ExampleRequestAttributes `json:"attributes" validate:"required"`
}

// ExampleRequestAttributes defines the subset of Example domain model attributes that are accepted
// for input data request binding
type ExampleRequestAttributes struct {
	Description *string `json:"description"`
	Status      *uint32 `json:"status"`
	Title       string  `json:"title"`
}

// Validate validates a Notification instance
func (e ExampleRequestAttributes) Validate() error {
	if err := v.ValidateStruct(&e,
		v.Field(&e.Title, v.Required),
	); err != nil {
		return err
	}

	return nil
}

// ------------------------------------------------------------------------------------------------

// ExampleResponseResource
type ExampleResponseResource struct {
	Type       string                   `json:"type"`
	ID         uuid.UUID                `json:"id"`
	Meta       *ExampleResourceMetadata `json:"meta,omitempty"`
	Attributes ExampleObjectAttributes  `json:"attributes"`
}

// ExampleResourceMetadata
type ExampleResourceMetadata struct{}

// ------------------------------------------------------------------------------------------------

// ExampleDomainModel an Example domain model that contains one or more ExampleObject(s)
// and related metadata
type ExampleDomainModel struct {
	Data []ExampleObject
	Meta *ModelMetadata
	Solo bool
}

type ModelMetadata struct {
	Paging pagination.PageMetadata
}

// ExampleObject
type ExampleObject struct {
	Attributes ExampleObjectAttributes
	Meta       any
	Related    any
}

// Example defines an Example domain model for application logic
type ExampleObjectAttributes struct {
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

func (m *ExampleDomainModel) FormatResponse() (*Response, error) {
	if m.Solo {
		resource := formatResource(&m.Data[0])
		response := &Response{Data: resource}
		return response, nil
	}

	meta := &ResponseMetadata{
		Page: pagination.PageMetadata{
			Limit:  m.Meta.Paging.Limit,
			Offset: m.Meta.Paging.Offset,
			Total:  m.Meta.Paging.Total,
		},
	}

	data := make([]ExampleResponseResource, 0, len(m.Data))
	for _, domo := range m.Data {
		resource := formatResource(&domo)
		data = append(data, resource)
	}
	response := &Response{
		Meta: meta,
		Data: data,
	}

	return response, nil
}

// serializeResource
func formatResource(domo *ExampleObject) ExampleResponseResource {
	return ExampleResponseResource{
		Type: "example", // TODO
		ID:   domo.Attributes.ID,
		// Meta: domo.Meta,
		Attributes: ExampleObjectAttributes{
			Title:       domo.Attributes.Title,
			Description: domo.Attributes.Description,
			Status:      domo.Attributes.Status,
			Enabled:     domo.Attributes.Enabled,
			Deleted:     domo.Attributes.Deleted,
			CreatedOn:   domo.Attributes.CreatedOn,
			CreatedBy:   domo.Attributes.CreatedBy,
			ModifiedOn:  domo.Attributes.ModifiedOn,
			ModifiedBy:  domo.Attributes.ModifiedBy,
		},
	}
}
