.PHONY: local dev prod

local:
	echo "Running in local mode..."
	go run ./cmd/auth/main.go --config=./config/local.yaml

dev:
	echo "Running in dev mode..."
	docker-compose -f docker-compose.dev.yaml up --build --remove-orphans


prod:
	echo "Running in prod mode..."
	docker-compose -f docker-compose.prod.yaml up --build --remove-orphans


run-linter:
	echo "Running linters..."
	golangci-lint run ./...