build:
	@go build -o ./bin/gobank ./cmd/main.go

run: build
	@./bin/gobank

test:
	@go test -v ./...