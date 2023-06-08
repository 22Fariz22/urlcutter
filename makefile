# ==============================================================================
# Main

run:
	go run cmd/cutter/main.go

run_db:
	 go run cmd/cutter/main.go -d="postgres://postgres:55555@127.0.0.1:5432/urlcutter"

test:
	go test -cover ./...

# ==============================================================================
# Go migrate postgresql

migrate_up:
	migrate -path migrations/ -database "postgresql://postgres:55555@localhost:5432/urlcutter?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations/ -database "postgresql://postgres:55555@localhost:5432/urlcutter?sslmode=disable" -verbose down
