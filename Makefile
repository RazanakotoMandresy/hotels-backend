build:
	@go build -C cmd/app -o ../../bin/hotels-backend .
run:build
	@./bin/hotels-backend
