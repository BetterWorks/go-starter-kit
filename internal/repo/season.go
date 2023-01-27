package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jasonsites/gosk-api/internal/types"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// SeasonRepoConfig defines the input to NewSeasonRepository
type SeasonRepoConfig struct {
	DBClient *pgxpool.Pool `validate:"required"`
	Logger   *types.Logger `validate:"required"`
}

// seasonRepository
type seasonRepository struct {
	db     *pgxpool.Pool
	logger *types.Logger
	name   string
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
		db:     c.DBClient,
		logger: logger,
		name:   "season",
	}

	return repo, nil
}

// Create
func (r *seasonRepository) Create(ctx context.Context, data any) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Season Repository Create called")

	requestData := data.(*types.SeasonRequestData)
	entity := types.SeasonEntity{}
	query := func() string {
		var (
			statement    = "INSERT INTO %s %s VALUES %s RETURNING %s"
			table        = r.name
			insertFields = "(created_by, deleted, description, enabled, status, title)"
			values       = "($1, $2, $3, $4, $5, $6)"
			returnFields = "created_by, deleted, description, enabled, id, status, title"
		)
		return fmt.Sprintf(statement, table, insertFields, values, returnFields)
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
		log.Error().Err(err)
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
	log.Info().Msg("Season Repository Delete called")

	entity := types.SeasonEntity{}
	query := func() string {
		var (
			statement    = "DELETE FROM %s WHERE id = ('%s'::uuid) RETURNING %s"
			table        = r.name
			returnFields = "created_by, deleted, description, enabled, id, modified_by, status, title"
		)
		return fmt.Sprintf(statement, table, id, returnFields)
	}()

	if err := r.db.QueryRow(ctx, query).Scan(
		&entity.CreatedBy,
		// &entity.CreatedOn,
		&entity.Deleted,
		&entity.Description,
		&entity.Enabled,
		&entity.ID,
		&entity.ModifiedBy,
		// &entity.ModifiedOn,
		&entity.Status,
		&entity.Title,
	); err != nil {
		log.Error().Err(err)
		return err
	}

	return nil
}

// Detail
func (r *seasonRepository) Detail(ctx context.Context, id uuid.UUID) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Season Repository Detail called")

	entity := types.SeasonEntity{}
	query := func() string {
		var (
			statement    = "SELECT %s FROM %s WHERE id = ('%s'::uuid)"
			returnFields = "created_by, deleted, description, enabled, id, modified_by, status, title"
			table        = r.name
		)
		return fmt.Sprintf(statement, returnFields, table, id)
	}()

	if err := r.db.QueryRow(ctx, query).Scan(
		&entity.CreatedBy,
		// &entity.CreatedOn,
		&entity.Deleted,
		&entity.Description,
		&entity.Enabled,
		&entity.ID,
		&entity.ModifiedBy,
		// &entity.ModifiedOn,
		&entity.Status,
		&entity.Title,
	); err != nil {
		log.Error().Err(err)
		return nil, err
	}

	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{
		Data: []types.RepoResultEntity{entityWrapper},
	}
	fmt.Printf("Result in seasonRepository.Create: %+v\n", result)

	return result, nil
}

// List
func (r *seasonRepository) List(ctx context.Context, m *types.ListMeta) ([]*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Season Repository List called")

	data := make([]*types.RepoResult, 2)
	return data, nil
}

// Update
func (r *seasonRepository) Update(ctx context.Context, data any, id uuid.UUID) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Season Repository Update called")

	requestData := data.(*types.SeasonRequestData)
	entity := types.SeasonEntity{}
	query := func() string {
		var (
			statement    = "UPDATE %s SET %s WHERE id = ('%s'::uuid) RETURNING %s"
			table        = r.name
			values       = "created_by=$1,deleted=$2,description=$3,enabled=$4,status=$5,title=$6"
			returnFields = "created_by, deleted, description, enabled, id, status, title"
		)
		return fmt.Sprintf(statement, table, values, id, returnFields)
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
		log.Error().Err(err)
		return nil, err
	}

	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{Data: []types.RepoResultEntity{entityWrapper}}

	return result, nil
}
