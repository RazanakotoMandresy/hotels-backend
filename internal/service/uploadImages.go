package service

import (
	"context"
	"errors"

	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/RazanakotoMandresy/hotels-backend/middleware"
)

func (s Service) UploadImages(ctx context.Context, hotelUUID string, file multipart.File, handler *multipart.FileHeader) (string, error) {
	hotels, err := s.GetHotel(ctx, hotelUUID)
	if err != nil {
		return "", err
	}
	userUUID := middleware.GetUserUUIDInAuth(ctx)
	if userUUID != hotels.CreatedBy {
		return "", errors.New("you are not the creator of this hotels")
	}
	if len(hotels.Images) > 8 {
		return "", errors.New("8 images per hotels maximum")
	}
	splitedName := strings.Split(handler.Filename, ".")
	destFile := "./uploads/" + splitedName[0] + hotelUUID + "." + splitedName[1]
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
	return "", nil
}
