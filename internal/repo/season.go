package repo

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jasonsites/gosk-api/internal/types"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// SeasonRepoConfig defines the input to NewSeasonRepository
type SeasonRepoConfig struct {
	Logger           *types.Logger `validate:"required"`
	PostgreSQLClient *sql.DB       `validate:"required"`
}

// seasonRepository
type seasonRepository struct {
	logger           *types.Logger
	postgreSQLClient *sql.DB
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
		logger:           logger,
		postgreSQLClient: c.PostgreSQLClient,
	}

	return repo, nil
}

// Create
func (r *seasonRepository) Create(data any) (*types.RepoResult, error) {
	log := r.logger.Log.With().Str("", "").Logger()
	log.Info().Msg("seasonRepository Create called")

	season := data.(*types.Season)
	season.ID = uuid.New() // mock ID return from DB
	entity := types.RepoResultEntity{Attributes: *season}

	result := &types.RepoResult{
		Data: []types.RepoResultEntity{entity},
	}
	fmt.Printf("Result in seasonRepository.Create: %+v\n", result)

	return result, nil
}

// Delete
func (r *seasonRepository) Delete(id uuid.UUID) error {
	fmt.Printf("ID in seasonRepository.Delete: %s\n", id)
	return nil
}

// Detail
func (r *seasonRepository) Detail(id uuid.UUID) (*types.RepoResult, error) {
	fmt.Printf("ID in seasonRepository.Detail: %s\n", id)

	data := &types.Season{
		ID: id,
	}
	entity := types.RepoResultEntity{Attributes: *data}

	result := &types.RepoResult{
		Data: []types.RepoResultEntity{entity},
	}
	fmt.Printf("Result in seasonRepository.Detail: %+v\n", result)
	return result, nil
}

// List
func (r *seasonRepository) List(m *types.ListMeta) ([]*types.RepoResult, error) {
	data := make([]*types.RepoResult, 2)
	return data, nil
}

// Update
func (r *seasonRepository) Update(data any) (*types.RepoResult, error) {
	season := data.(*types.Season)

	entity := types.RepoResultEntity{Attributes: *season}

	result := &types.RepoResult{
		Data: []types.RepoResultEntity{entity},
	}
	fmt.Printf("Result in seasonRepository.Update: %+v\n", result)

	return result, nil
}
