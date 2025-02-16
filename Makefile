include secrets/.env
export 

OAPI_CODEGEN_CONFIG_FILE=configs/oapi-codegen-config.yaml

all: clean setup run-migration gen-schemas gen-api run

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

clean:
	rm -rf internal/generated/*

tidy:
	go mod tidy
