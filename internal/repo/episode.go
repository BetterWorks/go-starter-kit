package repo

import (
	"database/sql"
	"fmt"

	"github.com/jasonsites/gosk-api/internal/application/domain"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// EpisodeEntity
type episodeEntity struct {
	Deleted     bool
	Description string
	Director    string
	Enabled     bool
	ID          string
	SeasonID    string
	Status      int
	Title       string
	Year        uint16
}

// EpisodeRepoConfig defines the input to NewEpisodeRepository
type EpisodeRepoConfig struct {
	DBClient *sql.DB        `validate:"required"`
	Logger   *domain.Logger `validate:"required"`
}

// episodeRepository
type episodeRepository struct {
	db     *sql.DB
	logger *domain.Logger
}

// NewEpisodeRepository
func NewEpisodeRepository(c *EpisodeRepoConfig) (*episodeRepository, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	log := c.Logger.Log.With().Str("tags", "repo,episode").Logger()
	logger := &domain.Logger{
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
func (r *episodeRepository) Create(data *domain.Episode) (*domain.RepoResult, error) {
	log := r.logger.Log.With().Str("req_id", "").Logger()
	log.Info().Msg("episodeRepository Create called")

	data.ID = "9999" // mock ID return from DB
	entity := domain.RepoResultEntity{Attributes: *data}

	result := &domain.RepoResult{
		Data: []domain.RepoResultEntity{entity},
	}
	fmt.Printf("Result in episodeRepository.Create: %+v\n", result)

	return result, nil
}

// Delete
func (r *episodeRepository) Delete(id string) error {
	log := r.logger.Log.With().Str("tags", "repo").Logger()
	log.Info().Msg("episodeRepository Delete called")

	fmt.Printf("ID in episodeRepository.Delete: %s\n", id)
	return nil
}

// Detail
func (r *episodeRepository) Detail(id string) (*domain.RepoResult, error) {
	log := r.logger.Log.With().Str("req_id", "").Logger()
	log.Info().Msg("episodeRepository Create called")

	fmt.Printf("ID in episodeRepository.Detail: %s\n", id)

	rows, err := r.db.Query("select * from episode where id = ?", id)
	if err != nil {
		log.Error().Err(err).Msg("error on episodeRepository.Detail db query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&episodeEntity{})
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
	// result := &domain.RepoResult{}

	data := &domain.Episode{
		ID: id,
	}
	entity := domain.RepoResultEntity{Attributes: *data}

	result := &domain.RepoResult{
		Data: []domain.RepoResultEntity{entity},
	}
	fmt.Printf("Result in episodeRepository.Detail: %+v\n", result)
	return result, nil
}

// List
func (r *episodeRepository) List(m *domain.ListMeta) ([]*domain.RepoResult, error) {
	data := make([]*domain.RepoResult, 2)
	return data, nil
}

// Update
func (r *episodeRepository) Update(data *domain.Episode) (*domain.RepoResult, error) {
	entity := domain.RepoResultEntity{Attributes: *data}

	result := &domain.RepoResult{
		Data: []domain.RepoResultEntity{entity},
	}
	fmt.Printf("Result in episodeRepository.Update: %+v\n", result)

	return result, nil
}
