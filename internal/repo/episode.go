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
	}

	return repo, nil
}

// Create
func (r *episodeRepository) Create(ctx context.Context, data any) (*types.RepoResult, error) {
	log := r.logger.Log.With().Str("req_id", "").Logger()
	log.Info().Msg("episodeRepository Create called")

	requestData := data.(*types.EpisodeRequestData)
	fmt.Printf("\n\nEPISODE REQUEST DATA: %+v\n\n", requestData)

	tableName := "episode"
	insertFields := "(created_by, deleted, description, director, enabled, season_id, status, title, year)"
	values := "($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	returnFields := "created_by, deleted, description, director, enabled, id, season_id, status, title, year"
	query := fmt.Sprintf("INSERT INTO %s %s VALUES %s RETURNING %s", tableName, insertFields, values, returnFields)

	entity := types.EpisodeEntity{}
	if err := r.db.QueryRow(
		context.Background(),
		query,
		9999,
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
		return nil, err
	}
	fmt.Printf("\n\nEPISODE RETURNED FROM DB: %+v\n\n", entity)

	entityWrapper := types.RepoResultEntity{Attributes: entity}

	result := &types.RepoResult{
		Data: []types.RepoResultEntity{entityWrapper},
	}
	fmt.Printf("Result in episodeRepository.Create: %+v\n", result)

	return result, nil
}

// Delete
func (r *episodeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	log := r.logger.Log.With().Str("tags", "repo").Logger()
	log.Info().Msg("episodeRepository Delete called")

	fmt.Printf("ID in episodeRepository.Delete: %s\n", id)
	return nil
}

// Detail
func (r *episodeRepository) Detail(ctx context.Context, id uuid.UUID) (*types.RepoResult, error) {
	log := r.logger.Log.With().Str("req_id", "").Logger()
	log.Info().Msg("episodeRepository Create called")

	fmt.Printf("ID in episodeRepository.Detail: %s\n", id)

	rows, err := r.db.Query(ctx, "SELECT * FROM episode WHERE id = ?", id)
	if err != nil {
		log.Error().Err(err).Msg("error on episodeRepository.Detail db query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&types.EpisodeEntity{})
		if err != nil {
			log.Error().Err(err).Msg("error on episodeRepository.Detail row scan")
		}
		fmt.Printf("rows in episodeRepository.Detail: %+v\n", rows)
	}
	err = rows.Err()
	if err != nil {
		log.Error().Err(err).Msg("rows error in episodeRepository.Detail")
	}

	// // TODO: marshal rows to RepoResult
	// result := &types.RepoResult{}

	data := &types.Episode{
		ID: id,
	}
	entity := types.RepoResultEntity{Attributes: *data}

	result := &types.RepoResult{
		Data: []types.RepoResultEntity{entity},
	}
	fmt.Printf("Result in episodeRepository.Detail: %+v\n", result)
	return result, nil
}

// List
func (r *episodeRepository) List(ctx context.Context, m *types.ListMeta) ([]*types.RepoResult, error) {
	data := make([]*types.RepoResult, 2)
	return data, nil
}

// Update
func (r *episodeRepository) Update(ctx context.Context, data any, id uuid.UUID) (*types.RepoResult, error) {
	episode := data.(*types.Episode)

	entity := types.RepoResultEntity{Attributes: *episode}

	result := &types.RepoResult{
		Data: []types.RepoResultEntity{entity},
	}
	fmt.Printf("Result in episodeRepository.Update: %+v\n", result)

	return result, nil
}
