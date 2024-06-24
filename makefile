build:
	@go build -o ./bin/gobank ./cmd/main.go

run: build
	@./bin/gobank

seed: build
	@./bin/gobank --seed

test:
	@go test ./... -v

db-dev:
	@docker-compose -f ./deployments/docker-compose.yaml up -d

buildImageAmd64:
	@docker build --platform=linux/amd64 -t myapp .

buildImageARM:
	@docker build --platform=linux/amd64 -t myapp .

pushImageAmd: buildImageAmd64
	docker push fernandounuts/gobank-amd64

upProd:
	@docker compose up --build