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
	}

	return repo, nil
}

// Create
func (r *seasonRepository) Create(ctx context.Context, data any) (*types.RepoResult, error) {
	log := r.logger.Log.With().Str("req_id", "").Logger()
	log.Info().Msg("seasonRepository Create called")

	requestData := data.(*types.SeasonRequestData)

	tableName := "season"
	insertFields := "(created_by, deleted, description, enabled, status, title)"
	values := "($1, $2, $3, $4, $5, $6)"
	returnFields := "created_by, deleted, description, enabled, id, status, title"
	query := fmt.Sprintf("INSERT INTO %s %s VALUES %s RETURNING %s", tableName, insertFields, values, returnFields)

	entity := types.SeasonEntity{}
	if err := r.db.QueryRow(
		ctx,
		query,
		9999,
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
		return nil, err
	}

	entityWrapper := types.RepoResultEntity{Attributes: entity}

	result := &types.RepoResult{
		Data: []types.RepoResultEntity{entityWrapper},
	}
	fmt.Printf("Result in seasonRepository.Create: %+v\n", result)

	return result, nil
}

// Delete
func (r *seasonRepository) Delete(ctx context.Context, id uuid.UUID) error {
	fmt.Printf("ID in seasonRepository.Delete: %s\n", id)
	return nil
}

// Detail
func (r *seasonRepository) Detail(ctx context.Context, id uuid.UUID) (*types.RepoResult, error) {
	fmt.Printf("ID in seasonRepository.Detail: %s\n", id)

	tableName := "season"
	returnFields := "created_by, deleted, description, enabled, id, modified_by, status, title"
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ('%s'::uuid)", returnFields, tableName, id)

	entity := types.SeasonEntity{}
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
	data := make([]*types.RepoResult, 2)
	return data, nil
}

// Update
func (r *seasonRepository) Update(ctx context.Context, data any) (*types.RepoResult, error) {
	season := data.(*types.Season)

	entity := types.RepoResultEntity{Attributes: *season}

	result := &types.RepoResult{
		Data: []types.RepoResultEntity{entity},
	}
	fmt.Printf("Result in seasonRepository.Update: %+v\n", result)

	return result, nil
}
