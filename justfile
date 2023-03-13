set shell := ["bash", "-uc"]

# Defaults ========================================================================================
project := 'domain'

# Commands ========================================================================================
# show this help
help:
  just --list

# remove build related files
clean:
  rm -rf bin
  rm -rf out
  rm -f profile.cov

# Lint ============================================================================================
# lint all
# lint: TODO

# Migrations ======================================================================================
# migrate down
migrate-down db +step='-all':
  migrate -path ./database/migrations -database postgres://postgres:postgres@domain_db:5432/{{db}}?sslmode=disable down {{step}}

# migrate up
migrate-up db *step:
  migrate -path ./database/migrations -database postgres://postgres:postgres@domain_db:5432/{{db}}?sslmode=disable up {{step}}

# migrate up -all (alias)
migrate:
  just migrate-up svcdb

# create migration with {{name}}
migrate-create name:
	migrate create -ext sql -dir ./database/migrations -format unix {{name}}

# Run =============================================================================================
# run http server in dev mode
serve-dev:
  just migrate
  go run ./cmd/httpserver/main.go

# run http server in dev mode with air monitor
serve-air:
  just migrate
  air

# Test ============================================================================================
# run tests
test:
  just migrate-up testdb
  go test -v ./...

# run tests with coverage report
coverage:
  just migrate-up testdb
  # go test -v ./test/integration/resource
  gotestsum --jsonfile ./test/coverage/coverage.log -- -covermode=count -coverprofile=./test/coverage/profile.cov ./test/integration/resource

# html coverage report
covreport:
  go tool cover -html=./coverage/profile.cov
