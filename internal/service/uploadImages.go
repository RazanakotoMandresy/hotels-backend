package service

import (
	"context"
	"fmt"
	"math/rand"

	// "crypto/rand"
	"io"
	"net/http"
	"os"
	"strings"
)

func (s Service) UploadImages(ctx context.Context, hotelUUID string, r *http.Request) (string, error) {
	hotels, err := s.GetHotel(ctx, hotelUUID)
	if err != nil {
		return "", err
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()
	splitedName := strings.Split(handler.Filename, ".")
	newFilename := splitedName[0] + hotelUUID + fmt.Sprint(rand.Int()) + "." + splitedName[1]
	f, err := os.OpenFile(newFilename, os.O_WRONLY|os.O_CREATE, 06666)
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
	return newFilename, nil
}
