package repo

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jasonsites/gosk-api/internal/application/domain"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// seasonEntity
type seasonEntity struct {
	Deleted     bool
	Description string
	Enabled     bool
	ID          string
	Status      int
	Title       string
}

// SeasonRepoConfig defines the input to NewSeasonRepository
type SeasonRepoConfig struct {
	Logger           *domain.Logger `validate:"required"`
	PostgreSQLClient *sql.DB        `validate:"required"`
}

// seasonRepository
type seasonRepository struct {
	logger           *domain.Logger
	postgreSQLClient *sql.DB
}

// NewSeasonRepository
func NewSeasonRepository(c *SeasonRepoConfig) (*seasonRepository, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	log := c.Logger.Log.With().Str("tags", "repo,season").Logger()
	logger := &domain.Logger{
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
func (r *seasonRepository) Create(data any) (*domain.RepoResult, error) {
	log := r.logger.Log.With().Str("", "").Logger()
	log.Info().Msg("seasonRepository Create called")

	season := data.(*domain.Season)
	season.ID = uuid.New() // mock ID return from DB
	entity := domain.RepoResultEntity{Attributes: *season}

	result := &domain.RepoResult{
		Data: []domain.RepoResultEntity{entity},
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
func (r *seasonRepository) Detail(id uuid.UUID) (*domain.RepoResult, error) {
	fmt.Printf("ID in seasonRepository.Detail: %s\n", id)

	data := &domain.Season{
		ID: id,
	}
	entity := domain.RepoResultEntity{Attributes: *data}

	result := &domain.RepoResult{
		Data: []domain.RepoResultEntity{entity},
	}
	fmt.Printf("Result in seasonRepository.Detail: %+v\n", result)
	return result, nil
}

// List
func (r *seasonRepository) List(m *domain.ListMeta) ([]*domain.RepoResult, error) {
	data := make([]*domain.RepoResult, 2)
	return data, nil
}

// Update
func (r *seasonRepository) Update(data any) (*domain.RepoResult, error) {
	season := data.(*domain.Season)

	entity := domain.RepoResultEntity{Attributes: *season}

	result := &domain.RepoResult{
		Data: []domain.RepoResultEntity{entity},
	}
	fmt.Printf("Result in seasonRepository.Update: %+v\n", result)

	return result, nil
}
