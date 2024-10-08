package main

import (
	"context"
	"log"

	"github.com/RazanakotoMandresy/hotels-backend/http/rest/handler"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run(ctx context.Context) error {
	server, err := handler.NewServer()
	if err != nil {
		return err
	}
	err = server.Run(ctx)
	return err
}
