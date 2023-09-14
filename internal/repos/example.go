package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/BetterWorks/gosk-api/internal/core/cerror"
	"github.com/BetterWorks/gosk-api/internal/core/entities"
	"github.com/BetterWorks/gosk-api/internal/core/interfaces"
	"github.com/BetterWorks/gosk-api/internal/core/logger"
	"github.com/BetterWorks/gosk-api/internal/core/models"
	"github.com/BetterWorks/gosk-api/internal/core/pagination"
	"github.com/BetterWorks/gosk-api/internal/core/query"
	"github.com/BetterWorks/gosk-api/internal/core/trace"
	"github.com/BetterWorks/gosk-api/internal/core/validation"
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
	DBClient *pgxpool.Pool  `validate:"required"`
	Logger   *logger.Logger `validate:"required"`
}

// exampleRepository
type exampleRepository struct {
	Entity exampleEntityDefinition
	db     *pgxpool.Pool
	logger *logger.Logger
}

// NewExampleRepository returns a new exampleRepository instance
func NewExampleRepository(c *ExampleRepoConfig) (*exampleRepository, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	log := c.Logger.Log.With().Str("tags", "repo,example").Logger()
	logger := &logger.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	repo := &exampleRepository{
		Entity: exampleEntity,
		db:     c.DBClient,
		logger: logger,
	}

	return repo, nil
}

// Create
func (r *exampleRepository) Create(ctx context.Context, data *models.ExampleInputData) (interfaces.DomainModel, error) {
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
			field.Deleted,
			field.Description,
			field.Enabled,
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
		requestData.Deleted,
		description,
		requestData.Enabled,
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
		log.Error().Err(err).Send()
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
		log.Error().Err(err).Send()
		return err
	}

	return nil
}

// Detail
func (r *exampleRepository) Detail(ctx context.Context, id uuid.UUID) (interfaces.DomainModel, error) {
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
	if scanErr := r.db.QueryRow(ctx, query).Scan(
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
	); scanErr != nil {
		log.Error().Err(scanErr).Send()
		err := cerror.NewNotFoundError(
			fmt.Sprintf("unable to find %s with id '%s'", r.Entity.Name, id),
		)
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
func (r *exampleRepository) List(ctx context.Context, q query.QueryData) (interfaces.DomainModel, error) {
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
		log.Error().Err(err).Send()
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
			log.Error().Err(err).Send()
			return nil, err
		}

		model.Data = append(model.Data, entity)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	// TODO: Investigate https://stackoverflow.com/questions/28888375/run-a-query-with-a-limit-offset-and-also-get-the-total-number-of-rows
	// query for total count
	var total int
	totalQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", r.Entity.Name)
	if err := r.db.QueryRow(ctx, totalQuery).Scan(&total); err != nil {
		log.Error().Err(err).Send()
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
func (r *exampleRepository) Update(ctx context.Context, data *models.ExampleInputData, id uuid.UUID) (interfaces.DomainModel, error) {
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
			field.Deleted,
			field.Description,
			field.Enabled,
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
		requestData.Deleted,
		description,
		requestData.Enabled,
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
		log.Error().Err(err).Send()
		return nil, err
	}

	model := &entities.ExampleEntityModel{
		Data: []entities.ExampleEntity{entity},
		Solo: true,
	}

	result := model.Unmarshal()

	return result, nil
}
