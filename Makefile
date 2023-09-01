include .env

db_up:
	docker compose up -d

db_down:
	docker compose down

m_up:
	migrate -path db/migrations -database "${POSTGRES_URL}&sslmode=disable" -verbose up

m_down:
	migrate -path db/migrations -database "${POSTGRES_URL}&sslmode=disable" -verbose down

m_create:
	migrate create -ext sql -dir db/migrations $(m_name)

api:
	go run ./cmd/api.go

