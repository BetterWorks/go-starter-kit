package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jasonsites/gosk-api/internal/types"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// seasonEntityDefinition
type seasonEntityDefinition struct {
	Field seasonEntityFields
	Name  string
}

// seasonEntityFields
type seasonEntityFields struct {
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

// seasonEntity
var seasonEntity = seasonEntityDefinition{
	Name: "season",
	Field: seasonEntityFields{
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

// SeasonRepoConfig defines the input to NewSeasonRepository
type SeasonRepoConfig struct {
	DBClient *pgxpool.Pool `validate:"required"`
	Logger   *types.Logger `validate:"required"`
}

// seasonRepository
type seasonRepository struct {
	Entity seasonEntityDefinition
	db     *pgxpool.Pool
	logger *types.Logger
}

// NewSeasonRepository
func NewSeasonRepository(c *SeasonRepoConfig) (*seasonRepository, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	log := c.Logger.Log.With().Str("tags", "repo,season").Logger()
	logger := &types.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	repo := &seasonRepository{
		Entity: seasonEntity,
		db:     c.DBClient,
		logger: logger,
	}

	return repo, nil
}

// Create
func (r *seasonRepository) Create(ctx context.Context, data any) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()

	requestData := data.(*types.SeasonRequestData)
	entity := types.SeasonEntity{}

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
			field.Deleted,
			field.Description,
			field.Enabled,
			field.ID,
			field.Status,
			field.Title,
		)

		return fmt.Sprintf(statement, name, insertFields, values, returnFields)
	}()

	if err := r.db.QueryRow(
		ctx,
		query,
		9999, // TODO: temp mock for `created_by` id
		requestData.Deleted,
		requestData.Description,
		requestData.Enabled,
		requestData.Status,
		requestData.Title,
	).Scan(
		&entity.CreatedBy,
		&entity.Deleted,
		&entity.Description,
		&entity.Enabled,
		&entity.ID,
		&entity.Status,
		&entity.Title,
	); err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{Data: []types.RepoResultEntity{entityWrapper}}

	return result, nil
}

// Delete
func (r *seasonRepository) Delete(ctx context.Context, id uuid.UUID) error {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()

	entity := types.SeasonEntity{}
	query := func() string {
		var (
			statement = "DELETE FROM %s WHERE id = ('%s'::uuid) RETURNING %s"
			field     = r.Entity.Field
			name      = r.Entity.Name
		)

		returnFields := buildReturnFields(
			field.CreatedBy,
			field.Deleted,
			field.Description,
			field.Enabled,
			field.ID,
			field.ModifiedBy,
			field.Status,
			field.Title,
		)

		return fmt.Sprintf(statement, name, id, returnFields)
	}()

	if err := r.db.QueryRow(ctx, query).Scan(&entity.ID); err != nil {
		log.Error().Err(err).Msg("")
		return err
	}

	return nil
}

// Detail
func (r *seasonRepository) Detail(ctx context.Context, id uuid.UUID) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()

	entity := types.SeasonEntity{}
	query := func() string {
		var (
			statement = "SELECT %s FROM %s WHERE id = ('%s'::uuid)"
			field     = r.Entity.Field
			name      = r.Entity.Name
		)

		returnFields := buildReturnFields(
			field.CreatedBy,
			field.Deleted,
			field.Description,
			field.Enabled,
			field.ID,
			field.ModifiedBy,
			field.Status,
			field.Title,
		)

		return fmt.Sprintf(statement, returnFields, name, id)
	}()

	if err := r.db.QueryRow(ctx, query).Scan(
		&entity.CreatedBy,
		&entity.Deleted,
		&entity.Description,
		&entity.Enabled,
		&entity.ID,
		&entity.ModifiedBy,
		&entity.Status,
		&entity.Title,
	); err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{Data: []types.RepoResultEntity{entityWrapper}}

	return result, nil
}

// List
func (r *seasonRepository) List(ctx context.Context, m types.ListMeta) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()

	query := func() string {
		var (
			statement = "SELECT %s FROM %s ORDER BY %s %s LIMIT %s OFFSET %s"
			field     = r.Entity.Field
			name      = r.Entity.Name
			orderBy   = "title"
			direction = "asc"
			limit     = "2"
			offset    = "0"
		)

		returnFields := buildReturnFields(
			field.CreatedBy,
			field.Deleted,
			field.Description,
			field.Enabled,
			field.ID,
			field.ModifiedBy,
			field.Status,
			field.Title,
		)
		return fmt.Sprintf(statement, returnFields, name, orderBy, direction, limit, offset)
	}()

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	defer rows.Close()

	result := &types.RepoResult{
		Data: make([]types.RepoResultEntity, 0),
	}

	for rows.Next() {
		entity := types.SeasonEntity{}

		if err := rows.Scan(
			&entity.CreatedBy,
			&entity.Deleted,
			&entity.Description,
			&entity.Enabled,
			&entity.ID,
			&entity.ModifiedBy,
			&entity.Status,
			&entity.Title,
		); err != nil {
			log.Error().Err(err).Msg("")
			return nil, err
		}

		entityWrapper := types.RepoResultEntity{Attributes: entity}
		result.Data = append(result.Data, entityWrapper)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	return result, nil
}

// Update
func (r *seasonRepository) Update(ctx context.Context, data any, id uuid.UUID) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()

	requestData := data.(*types.SeasonRequestData)
	entity := types.SeasonEntity{}
	query := func() string {
		var (
			statement = "UPDATE %s SET %s WHERE id = ('%s'::uuid) RETURNING %s"
			name      = r.Entity.Name
			field     = r.Entity.Field
		)

		values := buildUpdateValues(
			field.CreatedBy,
			field.Deleted,
			field.Description,
			field.Enabled,
			field.Status,
			field.Title,
		)

		returnFields := buildReturnFields(
			field.CreatedBy,
			field.Deleted,
			field.Description,
			field.Enabled,
			field.ID,
			field.Status,
			field.Title,
		)

		return fmt.Sprintf(statement, name, values, id, returnFields)
	}()

	if err := r.db.QueryRow(
		ctx,
		query,
		9999, // TODO: temp mock for user id
		requestData.Deleted,
		requestData.Description,
		requestData.Enabled,
		requestData.Status,
		requestData.Title,
	).Scan(
		&entity.CreatedBy,
		&entity.Deleted,
		&entity.Description,
		&entity.Enabled,
		&entity.ID,
		&entity.Status,
		&entity.Title,
	); err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{Data: []types.RepoResultEntity{entityWrapper}}

	return result, nil
}
