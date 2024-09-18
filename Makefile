build:
	@go build -o bin/deliveryapp-backend .
run:build
	@./bin/deliveryapp-backend
