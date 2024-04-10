package exampletest

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/BetterWorks/go-starter-kit/internal/core/entities"
	"github.com/BetterWorks/go-starter-kit/test/testutils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
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

func pollUntilServerStartup(t *testing.T, timeoutSeconds int, intervalMilliseconds int) {
	t.Log("polling until health check passes")
	assert.Eventually(t, func() bool {
		req := testutils.SetRequestData("GET", "/domain/health", nil, nil)
		if res, err := http.DefaultClient.Do(req); err != nil {
			t.Log(err)
		} else if res.StatusCode == 200 {
			return true
		}
		return false
	}, time.Duration(timeoutSeconds)*time.Second, time.Duration(intervalMilliseconds)*time.Millisecond)
}
