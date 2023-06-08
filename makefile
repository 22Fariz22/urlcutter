# ==============================================================================
# Main

run:
	go run ./cmd/passbook/main.go

test:
	go test -cover ./...

# ==============================================================================
# Go migrate postgresql

migrate_up:
	migrate -path migrations/ -database "postgresql://postgres:55555@localhost:5432/urlcutter?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations/ -database "postgresql://postgres:55555@localhost:5432/urlcutter?sslmode=disable" -verbose down
