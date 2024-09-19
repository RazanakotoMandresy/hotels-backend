build:
	@go build -C cmd/app -o bin/deliveryapp-backend .
run:build
	@./bin/deliveryapp-backend
