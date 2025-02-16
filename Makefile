include secrets/.env
export 

OAPI_CODEGEN_CONFIG_FILE=configs/oapi-codegen-config.yaml

all: setup run-migration gen-schemas gen-api test run

run:
	go run cmd/app/main.go

setup:
	bash scripts/create_database.sh

run-migration:
	bash scripts/run_migrations.sh

gen-schemas:
	go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgresql://$(PGUSER):$(PGPASSWORD)@$(PGHOST):$(PGPORT)/merch_store?sslmode=disable -schema=merch_store -path=./internal/generated

gen-api:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config $(OAPI_CODEGEN_CONFIG_FILE) api/schema.yaml

gen-mocks:
	go run github.com/golang/mock/mockgen -source=internal/services/auth/service.go --destination=mocks/services/auth/service.go
	go run github.com/golang/mock/mockgen -source=internal/services/jwt/service.go --destination=mocks/services/jwt/service.go
	go run github.com/golang/mock/mockgen -source=internal/services/products/service.go --destination=mocks/services/products/service.go
	go run github.com/golang/mock/mockgen -source=internal/services/storage/service.go --destination=mocks/services/storage/service.go

test:
	go test ./internal/handlers/...

clean:
	rm -rf internal/generated/*

tidy:
	go mod tidy
