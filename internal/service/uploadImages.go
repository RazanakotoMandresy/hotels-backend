package service

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

func (s Service) UploadImages(ctx context.Context, hotelUUID string, r *http.Request) (string, error) {
	hotels, err := s.GetHotel(ctx, hotelUUID)
	if err != nil {
		return "", err
	}
	// max memory
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()
	newFilename := handler.Filename + string(rand.Int())
	fmt.Println(newFilename)
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 06666)
	if err != nil {
		return "", err
	}	
	if _, err := io.Copy(f, file); err != nil {
		return "", err
	}
	hotels.Images = append(hotels.Images, handler.Filename)
	if err := s.repo.Update(ctx, *hotels); err != nil {
		return "", err
	}
	return handler.Filename, nil
}
