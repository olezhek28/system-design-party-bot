include .env
export

.PHONY: test-coverage
test-coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

.PHONY: update-all-deps
update-all-deps:
	go get -u ./... && go mod tidy

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN=$(DB_DSN_STG)

.PHONY: install-goose
.install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

.PHONY: compose-up
compose-up:
	docker-compose -f docker-compose-stg.yml up -d

.PHONY: compose-down
compose-down:
	docker-compose -f docker-compose-stg.yml down