package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/BetterWorks/go-starter-kit/internal/core/app"
	"github.com/BetterWorks/go-starter-kit/internal/core/cerror"
	"github.com/BetterWorks/go-starter-kit/internal/core/entities"
	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/BetterWorks/go-starter-kit/internal/core/models"
	"github.com/BetterWorks/go-starter-kit/internal/core/pagination"
	"github.com/BetterWorks/go-starter-kit/internal/core/query"
	"github.com/BetterWorks/go-starter-kit/internal/core/trace"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// exampleEntityDefinition
type exampleEntityDefinition struct {
	Field exampleEntityFields
	Name  string
}

// exampleEntityFields
type exampleEntityFields struct {
	CreatedBy   string
	CreatedOn   string
	Deleted     string
	Description string
	Enabled     string
	ID          string
	ModifiedBy  string
	ModifiedOn  string
	Status      string
	Title       string
}

// exampleEntity
var exampleEntity = exampleEntityDefinition{
	Name: "example_entity",
	Field: exampleEntityFields{
		CreatedBy:   "created_by",
		CreatedOn:   "created_on",
		Deleted:     "deleted",
		Description: "description",
		Enabled:     "enabled",
		ID:          "id",
		ModifiedBy:  "modified_by",
		ModifiedOn:  "modified_on",
		Status:      "status",
		Title:       "title",
	},
}

// ExampleRepoConfig defines the input to NewExampleRepository
type ExampleRepoConfig struct {
	DBClient *pgxpool.Pool        `validate:"required"`
	Logger   *logger.CustomLogger `validate:"required"`
}

// exampleRepository
type exampleRepository struct {
	Entity exampleEntityDefinition
	db     *pgxpool.Pool
	logger *logger.CustomLogger
}

// NewExampleRepository returns a new exampleRepository instance
func NewExampleRepository(c *ExampleRepoConfig) (*exampleRepository, error) {
	if err := app.Validator.Validate.Struct(c); err != nil {
		return nil, err
	}

	repo := &exampleRepository{
		Entity: exampleEntity,
		db:     c.DBClient,
		logger: c.Logger,
	}

	return repo, nil
}

// Create
func (r *exampleRepository) Create(ctx context.Context, data *models.ExampleRequestAttributes) (*models.ExampleDomainModel, error) {
	traceID := trace.GetTraceIDFromContext(ctx)
	log := r.logger.CreateContextLogger(traceID)

	// build sql query
	query := func() string {
		var (
			statement = "INSERT INTO %s %s VALUES %s RETURNING %s"
			field     = r.Entity.Field
			name      = r.Entity.Name
		)

		insertFields, values := buildInsertFieldsAndValues(
			field.CreatedBy,
			field.Description,
			field.Status,
			field.Title,
		)

		returnFields := buildReturnFields(
			field.CreatedBy,
			field.CreatedOn,
			field.Deleted,
			field.Description,
			field.Enabled,
			field.ID,
			field.ModifiedBy,
			field.ModifiedOn,
			field.Status,
			field.Title,
		)

		return fmt.Sprintf(statement, name, insertFields, values, returnFields)
	}()

	// gather data from request, handling for nullable fields
	requestData := data

	var (
		createdBy   = 9999 // TODO: temp mock for user id
		description *string
		status      *uint32
	)

	if requestData.Description != nil {
		description = requestData.Description
	}
	if requestData.Status != nil {
		status = requestData.Status
	}

	// create new entity for db row scan and execute query
	entity := entities.ExampleEntity{}
	if err := r.db.QueryRow(
		ctx,
		query,
		createdBy,
		description,
		status,
		requestData.Title,
	).Scan(
		&entity.CreatedBy,
		&entity.CreatedOn,
		&entity.Deleted,
		&entity.Description,
		&entity.Enabled,
		&entity.ID,
		&entity.ModifiedBy,
		&entity.ModifiedOn,
		&entity.Status,
		&entity.Title,
	); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	model := &entities.ExampleEntityModel{
		Data: []entities.ExampleEntity{entity},
		Solo: true,
	}

	result := model.Unmarshal()

	return result, nil
}

// Delete
func (r *exampleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	traceID := trace.GetTraceIDFromContext(ctx)
	log := r.logger.CreateContextLogger(traceID)

	// build sql query
	query := func() string {
		var (
			statement = "DELETE FROM %s WHERE id = ('%s'::uuid) RETURNING %s"
			field     = r.Entity.Field
			name      = r.Entity.Name
		)

		returnFields := buildReturnFields(field.ID)

		return fmt.Sprintf(statement, name, id, returnFields)
	}()

	// create new entity for db row scan and execute query
	entity := entities.ExampleEntity{}
	if err := r.db.QueryRow(ctx, query).Scan(&entity.ID); err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

// Detail
func (r *exampleRepository) Detail(ctx context.Context, id uuid.UUID) (*models.ExampleDomainModel, error) {
	traceID := trace.GetTraceIDFromContext(ctx)
	log := r.logger.CreateContextLogger(traceID)

	// build sql query
	query := func() string {
		var (
			statement = "SELECT %s FROM %s WHERE id = ('%s'::uuid)"
			field     = r.Entity.Field
			name      = r.Entity.Name
		)

		returnFields := buildReturnFields(
			field.CreatedBy,
			field.CreatedOn,
			field.Deleted,
			field.Description,
			field.Enabled,
			field.ID,
			field.ModifiedBy,
			field.ModifiedOn,
			field.Status,
			field.Title,
		)

		return fmt.Sprintf(statement, returnFields, name, id)
	}()

	// create new entity for db row scan and execute query
	entity := entities.ExampleEntity{}
	if err := r.db.QueryRow(ctx, query).Scan(
		&entity.CreatedBy,
		&entity.CreatedOn,
		&entity.Deleted,
		&entity.Description,
		&entity.Enabled,
		&entity.ID,
		&entity.ModifiedBy,
		&entity.ModifiedOn,
		&entity.Status,
		&entity.Title,
	); err != nil {
		log.Error(err.Error())
		err := cerror.NewNotFoundError(nil, fmt.Sprintf("unable to find %s with id '%s'", r.Entity.Name, id))
		return nil, err
	}

	model := &entities.ExampleEntityModel{
		Data: []entities.ExampleEntity{entity},
		Solo: true,
	}

	result := model.Unmarshal()

	return result, nil
}

// List
func (r *exampleRepository) List(ctx context.Context, q query.QueryData) (*models.ExampleDomainModel, error) {
	traceID := trace.GetTraceIDFromContext(ctx)
	log := r.logger.CreateContextLogger(traceID)

	var (
		limit  = *q.Paging.Limit
		offset = *q.Paging.Offset
	)

	// build sql query
	query := func() string {
		var (
			statement  = "SELECT %s FROM %s ORDER BY %s %s LIMIT %s OFFSET %s"
			field      = r.Entity.Field
			name       = r.Entity.Name
			orderField = *q.Sorting.Attr
			orderDir   = *q.Sorting.Order
			limit      = fmt.Sprint(limit)
			offset     = fmt.Sprint(offset)
		)

		returnFields := buildReturnFields(
			field.CreatedBy,
			field.CreatedOn,
			field.Deleted,
			field.Description,
			field.Enabled,
			field.ID,
			field.ModifiedBy,
			field.ModifiedOn,
			field.Status,
			field.Title,
		)

		return fmt.Sprintf(statement, returnFields, name, orderField, orderDir, limit, offset)
	}()

	// execute query, returning rows
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	// create new entity model
	model := &entities.ExampleEntityModel{
		Data: make([]entities.ExampleEntity, 0, limit),
		Solo: false,
	}

	// scan row data into new entities, appending to repo result
	for rows.Next() {
		entity := entities.ExampleEntity{}

		if err := rows.Scan(
			&entity.CreatedBy,
			&entity.CreatedOn,
			&entity.Deleted,
			&entity.Description,
			&entity.Enabled,
			&entity.ID,
			&entity.ModifiedBy,
			&entity.ModifiedOn,
			&entity.Status,
			&entity.Title,
		); err != nil {
			log.Error(err.Error())
			return nil, err
		}

		model.Data = append(model.Data, entity)
	}

	if err := rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// TODO: Investigate https://stackoverflow.com/questions/28888375/run-a-query-with-a-limit-offset-and-also-get-the-total-number-of-rows
	// query for total count
	var total int
	totalQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", r.Entity.Name)
	if err := r.db.QueryRow(ctx, totalQuery).Scan(&total); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// populate repo result paging metadata
	model.Meta = &models.ModelMetadata{
		Paging: pagination.PageMetadata{
			Limit:  uint32(limit),
			Offset: uint32(offset),
			Total:  uint32(total),
		},
	}

	result := model.Unmarshal()

	return result, nil
}

// Update
func (r *exampleRepository) Update(ctx context.Context, data *models.ExampleRequestAttributes, id uuid.UUID) (*models.ExampleDomainModel, error) {
	traceID := trace.GetTraceIDFromContext(ctx)
	log := r.logger.CreateContextLogger(traceID)

	// build sql query
	query := func() string {
		var (
			statement = "UPDATE %s SET %s WHERE id = ('%s'::uuid) RETURNING %s"
			name      = r.Entity.Name
			field     = r.Entity.Field
		)

		values := buildUpdateValues(
			field.Description,
			field.ModifiedBy,
			field.ModifiedOn,
			field.Status,
			field.Title,
		)

		returnFields := buildReturnFields(
			field.CreatedBy,
			field.CreatedOn,
			field.Deleted,
			field.Description,
			field.Enabled,
			field.ID,
			field.ModifiedBy,
			field.ModifiedOn,
			field.Status,
			field.Title,
		)

		return fmt.Sprintf(statement, name, values, id, returnFields)
	}()

	// gather data from request, handling for nullable fields
	requestData := data

	var (
		description *string
		modifiedBy  = 9999 // TODO: temp mock for user id
		modifiedOn  = time.Now()
		status      *uint32
	)

	if requestData.Description != nil {
		description = requestData.Description
	}
	if requestData.Status != nil {
		status = requestData.Status
	}

	// create new entity for db row scan and execute query
	entity := entities.ExampleEntity{}
	if err := r.db.QueryRow(
		ctx,
		query,
		description,
		modifiedBy,
		modifiedOn,
		status,
		requestData.Title,
	).Scan(
		&entity.CreatedBy,
		&entity.CreatedOn,
		&entity.Deleted,
		&entity.Description,
		&entity.Enabled,
		&entity.ID,
		&entity.ModifiedBy,
		&entity.ModifiedOn,
		&entity.Status,
		&entity.Title,
	); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	model := &entities.ExampleEntityModel{
		Data: []entities.ExampleEntity{entity},
		Solo: true,
	}

	result := model.Unmarshal()

	return result, nil
}
