# ==============================================================================
# Main

run:
	go run cmd/app/main.go

run_db:
	 go run cmd/app/main.go -d="postgres://postgres:55555@127.0.0.1:5432/urlcutter"

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


# build app and pull postgres from docker
docker-build:
	docker build -t app:1.0 . && docker pull postgres

compose-up: ### Run docker-compose
	 docker-compose up --build -d postgres &&  docker-compose up --build -d app &&  docker-compose logs -f

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
