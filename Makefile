include .env

m_up:
	migrate -path db/migrations -database "${POSTGRES_URL}&sslmode=disable" -verbose up

m_down:
	migrate -path db/migrations -database "${POSTGRES_URL}&sslmode=disable" -verbose down

telegram:
	go run ./cmd/main/telegram.go

api:
	go run ./cmd/main/api.go