FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -C cmd/app -o ../../bin/hotels-backend .

EXPOSE 8080

CMD ["./bin/backend"]