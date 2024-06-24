build:
	@go build -o ./bin/gobank ./cmd/main.go

run: build
	@./bin/gobank

seed: build
	@./bin/gobank --seed

test:
	@go test ./... -v

db-up:
	@docker-compose -f ./deployments/docker-compose.yaml up -d
