include .env

.PHONY: tools
tools:
	go install github.com/automation-co/husky@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/conventionalcommit/commitlint@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/segmentio/golines@latest
	go install mvdan.cc/gofumpt@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install ariga.io/atlas/cmd/atlas@latest

.PHONY: hook
hook: tools
	husky install

.PHONY: update_dependencies
update_dependencies:
	go get -d -u -t ./...

.PHONY: install
install:
	go mod download

.PHONY: test
test:
	# go test -v -coverpkg=./... -coverprofile=profile.cov ./... -parallel 1 -failfast
	sh scripts/coverage.sh
	go tool cover -func profile.cov

DATABASE := "postgres://$(POSTGRESQL_USERNAME):$(POSTGRESQL_PASSWORD)@$(POSTGRESQL_HOST):$(POSTGRESQL_PORT)/$(POSTGRESQL_DATABASE)?sslmode=disable"

.PHONY: migrate_change
migrate_change:
	atlas migrate diff --env gorm change

.PHONY: migrate_up
migrate_up:
	@ if [ -z "$(step)" ]; then \
  		migrate -database ${DATABASE} -source file://infra/postgresql/migrations up; \
    else \
      migrate -database ${DATABASE} -source file://infra/postgresql/migrations up $(step); \
    fi

.PHONY: migrate_down
migrate_down:
	@ if [ -z "$(step)" ]; then \
  		migrate -database ${DATABASE} -source file://infra/postgresql/migrations down; \
    else \
      migrate -database ${DATABASE} -source file://infra/postgresql/migrations down $(step); \
    fi
