.PHONY: local dev dev_down prod prod_down

local:
	echo "Running in local mode..."
	go run ./cmd/auth/main.go --config=./config/local.yaml

dev:
	echo "Running in dev mode..."
	docker-compose -f docker-compose.dev.yaml up --build --remove-orphans

dev_down:
	docker-compose -f docker-compose.dev.yaml down

prod:
	echo "Running in prod mode..."
	docker-compose -f docker-compose.prod.yaml up --build --remove-orphans

prod_down:
	docker-compose -f docker-compose.prod.yaml down

run-linter:
	echo "Running linters..."
	golangci-lint run ./...

