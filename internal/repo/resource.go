package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/BetterWorks/gosk-api/internal/types"
	"github.com/BetterWorks/gosk-api/internal/validation"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// resourceEntityDefinition
type resourceEntityDefinition struct {
	Field resourceEntityFields
	Name  string
}

// resourceEntityFields
type resourceEntityFields struct {
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

// resourceEntity
var resourceEntity = resourceEntityDefinition{
	Name: "resource_entity",
	Field: resourceEntityFields{
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

// ResourceRepoConfig defines the input to NewResourceRepository
type ResourceRepoConfig struct {
	DBClient *pgxpool.Pool `validate:"required"`
	Logger   *types.Logger `validate:"required"`
}

// resourceRepository
type resourceRepository struct {
	Entity resourceEntityDefinition
	db     *pgxpool.Pool
	logger *types.Logger
}

// NewResourceRepository
func NewResourceRepository(c *ResourceRepoConfig) (*resourceRepository, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	log := c.Logger.Log.With().Str("tags", "repo,resource").Logger()
	logger := &types.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	repo := &resourceRepository{
		Entity: resourceEntity,
		db:     c.DBClient,
		logger: logger,
	}

	return repo, nil
}

// Create
func (r *resourceRepository) Create(ctx context.Context, data any) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()

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
	requestData := data.(*types.ResourceRequestData)

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
	entity := types.ResourceEntity{}
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

	// return new repo result from scanned entity
	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{Data: []types.RepoResultEntity{entityWrapper}}

	return result, nil
}

// Delete
func (r *resourceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()

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
	entity := types.ResourceEntity{}
	if err := r.db.QueryRow(ctx, query).Scan(&entity.ID); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}

// Detail
func (r *resourceRepository) Detail(ctx context.Context, id uuid.UUID) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()

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
	entity := types.ResourceEntity{}
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
		err := types.NewNotFoundError(
			fmt.Sprintf("unable to find %s with id '%s'", r.Entity.Name, id),
		)
		return nil, err
	}

	// return new repo result from scanned entity
	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{Data: []types.RepoResultEntity{entityWrapper}}

	return result, nil
}

// List
func (r *resourceRepository) List(ctx context.Context, q types.QueryData) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()

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
			orderField = *q.Sorting.Prop
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

	// create new repo result
	result := &types.RepoResult{
		Data: make([]types.RepoResultEntity, 0),
	}

	// scan row data into new entities, appending to repo result
	for rows.Next() {
		entity := types.ResourceEntity{}

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

		entityWrapper := types.RepoResultEntity{Attributes: entity}
		result.Data = append(result.Data, entityWrapper)
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
	result.Metadata = types.RepoResultMetadata{
		Paging: types.ListPaging{
			Limit:  uint32(limit),
			Offset: uint32(offset),
			Total:  uint32(total),
		},
	}

	return result, nil
}

// Update
func (r *resourceRepository) Update(ctx context.Context, data any, id uuid.UUID) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()

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
	requestData := data.(*types.ResourceRequestData)

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
	entity := types.ResourceEntity{}
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

	// return new repo result from scanned entity
	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{Data: []types.RepoResultEntity{entityWrapper}}

	return result, nil
}
