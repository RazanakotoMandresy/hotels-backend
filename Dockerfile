FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod tidy
ENV DATABASE_HOST = db  
ENV DATABASE_PORT = 5433
ENV DATABASE_USER = user
ENV DATABASE_PASSWORD = password
ENV DATABASE_NAME = hotels
ENV JWT_SECRET = secret
ENV CRYPT_KEY = CRYPT_KEY

RUN go build -C cmd/app -o ../../bin/hotels-backend .

EXPOSE 8080

CMD ["./bin/backend"]