run: 
	go run cmd/app/main.go

gen:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -package=api -generate types,spec api/schema.yaml > internal/generated/api.gen.go

tidy:
	go mod tidy

clean:
	rm -rf internal/generated/*
