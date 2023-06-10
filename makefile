# ==============================================================================
# Main

run:
	go run cmd/cutter/main.go -a ":5001"

run_db:
	 go run cmd/cutter/main.go -d="postgres://postgres:55555@127.0.0.1:5432/urlcutter"

test:
	go test -cover ./...

test-html:
	 go test -short -count=1 -race -coverprofile=coverage.out ./... &&  go tool cover -html=coverage.out

# ==============================================================================
# Go migrate postgresql

migrate_up:
	migrate -path migrations/ -database "postgresql://postgres:55555@localhost:5432/urlcutter?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations/ -database "postgresql://postgres:55555@localhost:5432/urlcutter?sslmode=disable" -verbose down

# ==============================================================================
# Docker compose commands

develop:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build
