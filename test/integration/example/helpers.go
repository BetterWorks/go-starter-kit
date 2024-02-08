package exampletest

import (
	"context"
	"fmt"

	"github.com/BetterWorks/go-starter-kit/internal/core/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

const routePrefix = "/domain/examples"

// insertRecord inserts a db record for use in test setup
func insertRecord(db *pgxpool.Pool) (*entities.ExampleEntity, error) {
	var (
		statement    = "INSERT INTO %s %s VALUES %s RETURNING id"
		name         = "example_entity"
		insertFields = "(created_by,deleted,description,enabled,status,title)"
		values       = "($1,$2,$3,$4,$5,$6)"
		query        = fmt.Sprintf(statement, name, insertFields, values)
	)

	var (
		createdBy   = 9999
		deleted     = false
		description = "test description"
		enabled     = true
		status      = 1
		title       = "test title"
	)

	// create new entity for db row scan and execute query
	entity := &entities.ExampleEntity{}
	if err := db.QueryRow(
		context.Background(),
		query,
		createdBy,
		deleted,
		description,
		enabled,
		status,
		title,
	).Scan(
		&entity.ID,
	); err != nil {
		return nil, err
	}

	return entity, nil
}
