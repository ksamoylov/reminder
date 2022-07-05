FROM golang:1.18-alpine

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go mod download && go mod verify
RUN go build -o reminder ./cmd/main/api.go
RUN go build -o reminder ./cmd/main/telegram.go

CMD ["/app/reminder"]