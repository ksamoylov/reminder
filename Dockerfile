FROM golang:1.18-alpine

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go mod download && go mod verify
RUN go build -o reminder ./cmd/main/api.go

CMD ["/app/reminder"]