# Load environment variables from .env
include .env
export

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	cd tests && go test -v -race -buildvcs

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## run: run the cmd/api application
.PHONY: run
run:
	go run cmd/app/main.go

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/api" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

## swagger/serve: run swagger-ui from openapi.yaml
.PHONY: swagger/serve
swagger/serve:
	docker run --detach --name swagger_doc --publish 80:8080 --env SWAGGER_JSON=/openapi.yaml --volume $$(pwd)/openapi.yaml:/openapi.yaml swaggerapi/swagger-ui:v5.10.5

## swagger/down: stop swagger-ui container
.PHONY: swagger/down
swagger/down:
	docker rm -f swagger_doc

# ==================================================================================== #
# SQL MIGRATIONS
# ==================================================================================== #

DB_DSN="mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_SCHEMA}"

## migrations/new name=$1: create a new database migration
.PHONY: migrations/new
migrations/new:
ifdef table
	migrate create -dir db/migrations -ext sql ${table}
else
	echo "must define \`table\` argument"
endif

## migrations/up: apply all up database migrations
.PHONY: migrations/up
migrations/up:
	migrate -path=db/migrations -database=${DB_DSN} up

## migrations/down: apply all down database migrations
.PHONY: migrations/down
migrations/down:
	migrate -path=db/migrations -database=${DB_DSN} down

## migrations/goto version=$1: migrate to a specific version number
.PHONY: migrations/goto
migrations/goto:
	migrate -path=db/migrations -database=${DB_DSN} goto ${version}

## migrations/force version=$1: force database migration
.PHONY: migrations/force
migrations/force:
	migrate -path=db/migrations -database=${DB_DSN} force ${version}
.PHONY: migrations/version

## migrations/version: print the current in-use migration version
migrations/version:
	migrate -path=db/migrations -database=${DB_DSN} version
