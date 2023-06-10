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


compose-up: ### Run docker-compose
	docker-compose up --build -d postgres rabbitmq && docker-compose logs -f

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
