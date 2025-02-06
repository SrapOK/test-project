tidy:
	go mod tidy

run:
	go run cmd/test-project/main.go

sqlc:
	sqlc generate

swag:
	swag fmt
	swag init -d="cmd/test-project,internal/app,internal/contract"