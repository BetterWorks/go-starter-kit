set shell := ["bash", "-uc"]

# Defaults ========================================================================================
project := 'domain'

# Commands ========================================================================================
# show this help
help:
  just --list

# remove build related files
clean:
  rm -rf out
  rm -f test/coverage

# Migrations ======================================================================================
# migrate down
migrate-down db +step='-all':
  migrate -path ./database/migrations -database postgres://postgres:postgres@postgres:5432/{{db}}?sslmode=disable down {{step}}

# migrate up
migrate-up db *step:
  migrate -path ./database/migrations -database postgres://postgres:postgres@postgres:5432/{{db}}?sslmode=disable up {{step}}

# migrate up -all (alias)
migrate:
  just migrate-up svcdb

# create migration with {{name}}
migrate-create name:
	migrate create -ext sql -dir ./database/migrations -format unix {{name}}

# Run =============================================================================================
# run {http}server in dev mode
serve-dev +protocol='http':
  go run ./cmd/{{protocol}}server/main.go

# run {http}server in dev mode with file monitor
serve +protocol='http':
  air --tmp_dir="out/tmp" --build.cmd="go build -mod readonly -o out/tmp/domain ./cmd/{{protocol}}server" --build.bin="out/tmp/domain"

# Test ============================================================================================
# run tests with {{pattern}} arguments
test +pattern='--format testname -- ./...':
  gotestsum {{pattern}}

# run integration tests (overridable with {{scope}} arguments)
test-int +scope='':
  just migrate-up testdb
  POSTGRES_DB=testdb just test --format testname -- -race ./test/integration/... {{scope}}

# run unit tests (overridable with {{scope}} arguments)
test-unit +scope='':
  just test --format testname -- -race ./internal/... {{scope}}

# run unit/inegration tests and generate coverage report, exclude packages from test folder from coverage
coverage:
  #!/usr/bin/env bash
  just migrate-up testdb
  packages=`echo $(go list ./... 2> /dev/null | grep -v /test/) | tr ' ' ","`
  POSTGRES_DB=testdb just test \
    --format testname \
    --jsonfile /app/test/coverage/json.log \
    --junitfile /app/test/coverage/junit.xml \
    -- -coverprofile=profile.cov -outputdir=test/coverage -coverpkg=$packages ./...
  just covreport

# html coverage report
covreport:
  go tool cover -html=test/coverage/profile.cov -o test/coverage/report.html

# generate openAPI docs
generate-openapi:
  rm -f ./openapi/*
  swag init -g ./cmd/httpserver/main.go --parseDependency --parseInternal -o ./openapi/ --outputTypes json
  curl -X 'POST' 'https://converter.swagger.io/api/convert' -H 'accept: application/json' -H 'Content-Type: application/json' -d '@openapi/swagger.json' > openapi/swagger_v3.json
