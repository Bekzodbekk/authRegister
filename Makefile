DB_URL = postgres://postgres:1@localhost:5432/postgres?sslmode=disable

migrate-up:
  migrate -path ./db/migrations -database ${DB_URL} up

migrate-down:
  migrate -path ./db/migrations -database ${DB_URL} down

migrate-force:
  migrate -path=db/migrations -database ${DB_URL} verbose -force 1
