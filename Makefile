run:
	go run cmd/ьain.go

build:
	go build -o bin/bookstore cmd.main.go

export POSTGRES_DB=postgres://student_n9:student_n9@localhost:5432/bookstore_n9?sslmode=disable

migrate-booksfile:
	migrate create -ext sql -dir pkg/migrations/ -seq books_table

migrate-authorsfile:
	migrate create -ext sql -dir pkg/migrations/ -seq authors_table

migrate-up:
	migrate -path pkg/migrations -database $(POSTGRES_DB) up

migrate-down:
	migrate -path pkg/migrations -database $(POSTGRES_DB) down

migrate-force:
	migrate -path pkg/migrations -database $(POSTGRES_DB) force $(version)

.PHONY: run build migrate-up migrate-down migrate-force
