include .env

m_up:
	migrate -path db/migrations -database "${POSTGRES_URL}&sslmode=disable" -verbose up

m_down:
	migrate -path db/migrations -database "${POSTGRES_URL}&sslmode=disable" -verbose down
