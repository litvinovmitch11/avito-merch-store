OAPI_CODEGEN_CONFIG_FILE = configs/oapi-codegen-config.yaml

all: clean gen run

run: 
	go run cmd/app/main.go

gen:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config $(OAPI_CODEGEN_CONFIG_FILE) api/schema.yaml

clean:
	rm -rf internal/generated/*

tidy:
	go mod tidy
