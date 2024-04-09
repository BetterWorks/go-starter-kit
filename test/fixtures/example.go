package fixtures

import (
	"time"

	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

type ExampleRequestAttributesBuilder struct {
	description string
	status      uint32
	title       string
}

func NewExampleRequestAttributesBuilder() *ExampleRequestAttributesBuilder {
	return &ExampleRequestAttributesBuilder{
		description: fake.Phrase(),
		status:      fake.Uint32(),
		title:       fake.Word(),
	}
}

func (b *ExampleRequestAttributesBuilder) Description(description string) *ExampleRequestAttributesBuilder {
	b.description = description
	return b
}

func (b *ExampleRequestAttributesBuilder) Status(status uint32) *ExampleRequestAttributesBuilder {
	b.status = status
	return b
}

func (b *ExampleRequestAttributesBuilder) Title(title string) *ExampleRequestAttributesBuilder {
	b.title = title
	return b
}

func (b *ExampleRequestAttributesBuilder) Build() *models.ExampleRequestAttributes {
	return &models.ExampleRequestAttributes{
		Description: &b.description,
		Status:      &b.status,
		Title:       b.title,
	}
}

type ExampleObjectAttributesBuilder struct {
	id          uuid.UUID
	title       string
	description *string
	status      *uint32
	enabled     bool
	deleted     bool
	createdOn   time.Time
	createdBy   uint32
	modifiedOn  *time.Time
	modifiedBy  *uint32
}

func NewExampleObjectAttributesBuilder() *ExampleObjectAttributesBuilder {
	description := fake.Word()
	status := fake.Uint32()
	modifiedOn := fake.Date()
	modifiedBy := fake.Uint32()
	return &ExampleObjectAttributesBuilder{
		id:          uuid.New(),
		title:       fake.Word(),
		description: &description,
		status:      &status,
		enabled:     fake.Bool(),
		deleted:     fake.Bool(),
		createdOn:   fake.Date(),
		createdBy:   fake.Uint32(),
		modifiedOn:  &modifiedOn,
		modifiedBy:  &modifiedBy,
	}
}

func (b *ExampleObjectAttributesBuilder) ID(id uuid.UUID) *ExampleObjectAttributesBuilder {
	b.id = id
	return b
}

func (b *ExampleObjectAttributesBuilder) Title(title string) *ExampleObjectAttributesBuilder {
	b.title = title
	return b
}

func (b *ExampleObjectAttributesBuilder) Description(description *string) *ExampleObjectAttributesBuilder {
	b.description = description
	return b
}

func (b *ExampleObjectAttributesBuilder) Status(status *uint32) *ExampleObjectAttributesBuilder {
	b.status = status
	return b
}

func (b *ExampleObjectAttributesBuilder) Enabled(enabled bool) *ExampleObjectAttributesBuilder {
	b.enabled = enabled
	return b
}

func (b *ExampleObjectAttributesBuilder) Deleted(deleted bool) *ExampleObjectAttributesBuilder {
	b.deleted = deleted
	return b
}

func (b *ExampleObjectAttributesBuilder) CreatedOn(createdOn time.Time) *ExampleObjectAttributesBuilder {
	b.createdOn = createdOn
	return b
}

func (b *ExampleObjectAttributesBuilder) CreatedBy(createdBy uint32) *ExampleObjectAttributesBuilder {
	b.createdBy = createdBy
	return b
}

func (b *ExampleObjectAttributesBuilder) ModifiedOn(modifiedOn *time.Time) *ExampleObjectAttributesBuilder {
	b.modifiedOn = modifiedOn
	return b
}

func (b *ExampleObjectAttributesBuilder) ModifiedBy(modifiedBy *uint32) *ExampleObjectAttributesBuilder {
	b.modifiedBy = modifiedBy
	return b
}

func (b *ExampleObjectAttributesBuilder) Build() *models.ExampleObjectAttributes {
	return &models.ExampleObjectAttributes{
		ID:          b.id,
		Title:       b.title,
		Description: b.description,
		Status:      b.status,
		Enabled:     b.enabled,
		Deleted:     b.deleted,
		CreatedOn:   b.createdOn,
		CreatedBy:   b.createdBy,
		ModifiedOn:  b.modifiedOn,
		ModifiedBy:  b.modifiedBy,
	}
}

type ExampleObjectBuilder struct {
	attributes models.ExampleObjectAttributes
	meta       any
	related    any
}

func NewExampleObjectBuilder() *ExampleObjectBuilder {
	return &ExampleObjectBuilder{
		attributes: *NewExampleObjectAttributesBuilder().Build(),
		meta:       nil,
		related:    nil,
	}
}

func (b *ExampleObjectBuilder) Attributes(attributes models.ExampleObjectAttributes) *ExampleObjectBuilder {
	b.attributes = attributes
	return b
}

func (b *ExampleObjectBuilder) Meta(meta any) *ExampleObjectBuilder {
	b.meta = meta
	return b
}

func (b *ExampleObjectBuilder) Related(related any) *ExampleObjectBuilder {
	b.related = related
	return b
}

func (b *ExampleObjectBuilder) Build() models.ExampleObject {
	return models.ExampleObject{
		Attributes: b.attributes,
		Meta:       b.meta,
		Related:    b.related,
	}
}
