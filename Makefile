OAPI_CODEGEN_CONFIG_FILE = configs/oapi-codegen-config.yaml

all: clean gen run

run:
	go run cmd/app/main.go

run-migration:
	bash scripts/run_migrations.sh

gen-api:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config $(OAPI_CODEGEN_CONFIG_FILE) api/schema.yaml

gen-schemas:
	go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgresql://postgres:postgres@localhost:5432/merch_store?sslmode=disable -schema=merch_store -path=./internal/generated

clean:
	rm -rf internal/generated/*

tidy:
	go mod tidy
