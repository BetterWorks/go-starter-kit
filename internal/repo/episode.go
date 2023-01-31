package repo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jasonsites/gosk-api/internal/types"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// EpisodeRepoConfig defines the input to NewEpisodeRepository
type EpisodeRepoConfig struct {
	DBClient *pgxpool.Pool `validate:"required"`
	Logger   *types.Logger `validate:"required"`
}

// episodeRepository
type episodeRepository struct {
	db     *pgxpool.Pool
	logger *types.Logger
	name   string
}

// NewEpisodeRepository
func NewEpisodeRepository(c *EpisodeRepoConfig) (*episodeRepository, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	log := c.Logger.Log.With().Str("tags", "repo,episode").Logger()
	logger := &types.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	repo := &episodeRepository{
		db:     c.DBClient,
		logger: logger,
		name:   "episode",
	}

	return repo, nil
}

// Create
func (r *episodeRepository) Create(ctx context.Context, data any) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Episode Repository Create called")

	requestData := data.(*types.EpisodeRequestData)
	entity := types.EpisodeEntity{}
	query := func() string {
		var (
			statement    = "INSERT INTO %s %s VALUES %s RETURNING %s"
			table        = r.name
			insertFields = "(created_by, deleted, description, director, enabled, season_id, status, title, year)"
			values       = "($1, $2, $3, $4, $5, $6, $7, $8, $9)"
			returnFields = "created_by, deleted, description, director, enabled, id, season_id, status, title, year"
		)
		return fmt.Sprintf(statement, table, insertFields, values, returnFields)
	}()

	if err := r.db.QueryRow(
		context.Background(),
		query,
		9999, // TODO
		requestData.Deleted,
		requestData.Description,
		requestData.Director,
		requestData.Enabled,
		requestData.SeasonID,
		requestData.Status,
		requestData.Title,
		requestData.Year,
	).Scan(
		&entity.CreatedBy,
		&entity.Deleted,
		&entity.Description,
		&entity.Director,
		&entity.Enabled,
		&entity.ID,
		&entity.SeasonID,
		&entity.Status,
		&entity.Title,
		&entity.Year,
	); err != nil {
		log.Error().Err(err)
		return nil, err
	}

	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{Data: []types.RepoResultEntity{entityWrapper}}

	return result, nil
}

// Delete
func (r *episodeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Episode Repository Delete called")

	entity := types.EpisodeEntity{}
	query := func() string {
		var (
			statement    = "DELETE FROM %s WHERE id = ('%s'::uuid) RETURNING %s"
			table        = r.name
			returnFields = "id"
		)
		return fmt.Sprintf(statement, table, id, returnFields)
	}()

	if err := r.db.QueryRow(ctx, query).Scan(&entity.ID); err != nil {
		log.Error().Err(err)
		return err
	}

	return nil
}

// Detail
func (r *episodeRepository) Detail(ctx context.Context, id uuid.UUID) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Episode Repository Detail called")

	entity := types.EpisodeEntity{}
	query := func() string {
		var (
			statement    = "SELECT %s FROM %s WHERE id = ('%s'::uuid)"
			returnFields = "created_by, deleted, description, director, enabled, id, season_id, status, title, year"
			table        = r.name
		)
		return fmt.Sprintf(statement, returnFields, table, id)
	}()

	if err := r.db.QueryRow(ctx, query).Scan(
		&entity.CreatedBy,
		&entity.Deleted,
		&entity.Description,
		&entity.Director,
		&entity.Enabled,
		&entity.ID,
		&entity.SeasonID,
		&entity.Status,
		&entity.Title,
		&entity.Year,
	); err != nil {
		log.Error().Err(err)
		return nil, err
	}

	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{Data: []types.RepoResultEntity{entityWrapper}}

	return result, nil
}

// List
func (r *episodeRepository) List(ctx context.Context, m types.ListMeta) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Episode Repository List called")

	query := func() string {
		var (
			statement    = "SELECT %s FROM %s ORDER BY %s %s LIMIT %s OFFSET %s"
			returnFields = "created_by, deleted, description, director, enabled, id, season_id, status, title, year"
			table        = r.name
			orderBy      = "id"
			direction    = "desc"
			limit        = "20"
			offset       = "0"
		)
		return fmt.Sprintf(statement, returnFields, table, orderBy, direction, limit, offset)
	}()

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	defer rows.Close()

	result := &types.RepoResult{
		Data: make([]types.RepoResultEntity, 0),
	}

	for rows.Next() {
		entity := types.EpisodeEntity{}

		if err := rows.Scan(
			&entity.CreatedBy,
			&entity.Deleted,
			&entity.Description,
			&entity.Director,
			&entity.Enabled,
			&entity.ID,
			&entity.SeasonID,
			&entity.Status,
			&entity.Title,
			&entity.Year,
		); err != nil {
			log.Error().Err(err)
			return nil, err
		}

		entityWrapper := types.RepoResultEntity{Attributes: entity}
		result.Data = append(result.Data, entityWrapper)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err)
		return nil, err
	}

	return result, nil
}

// Update
func (r *episodeRepository) Update(ctx context.Context, data any, id uuid.UUID) (*types.RepoResult, error) {
	requestId := ctx.Value(types.CorrelationContextKey).(*types.Trace).RequestID
	log := r.logger.Log.With().Str("req_id", requestId).Logger()
	log.Info().Msg("Episode Repository Update called")

	requestData := data.(*types.EpisodeRequestData)
	entity := types.EpisodeEntity{}
	query := func() string {
		var (
			statement    = "UPDATE %s SET %s WHERE id = ('%s'::uuid) RETURNING %s"
			table        = r.name
			values       = "created_by=$1,deleted=$2,description=$3,director=$4,enabled=$5,season_id=$6,status=$7,title=$8,year=$9"
			returnFields = "created_by, deleted, description, director, enabled, id, season_id, status, title, year"
		)
		return fmt.Sprintf(statement, table, values, id, returnFields)
	}()

	if err := r.db.QueryRow(
		ctx,
		query,
		9999, // TODO: temp mock for user id
		requestData.Deleted,
		requestData.Description,
		requestData.Director,
		requestData.Enabled,
		requestData.SeasonID,
		requestData.Status,
		requestData.Title,
		requestData.Year,
	).Scan(
		&entity.CreatedBy,
		&entity.Deleted,
		&entity.Description,
		&entity.Director,
		&entity.Enabled,
		&entity.ID,
		&entity.SeasonID,
		&entity.Status,
		&entity.Title,
		&entity.Year,
	); err != nil {
		log.Error().Err(err)
		return nil, err
	}

	entityWrapper := types.RepoResultEntity{Attributes: entity}
	result := &types.RepoResult{Data: []types.RepoResultEntity{entityWrapper}}

	return result, nil
}
