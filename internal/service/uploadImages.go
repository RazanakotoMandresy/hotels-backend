package service

import (
	"context"
	"fmt"
	"math/rand"

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
	destFile := "./uploads/" + splitedName[0] + hotelUUID + fmt.Sprint(rand.Int()) + "." + splitedName[1]
	out, err := os.Create(destFile)
	if err != nil {
		return "", err
	}
	defer out.Close()
	if _, err := io.Copy(out, file); err != nil {
		return "", err
	}
	hotels.Images = append(hotels.Images, handler.Filename)
	if err := s.repo.Update(ctx, *hotels); err != nil {
		return "", err
	}
	return destFile, nil
}
