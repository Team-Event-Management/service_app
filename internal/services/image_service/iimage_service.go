package imageservice

import (
	"context"
	"event_management/internal/models"
	"mime/multipart"

	"github.com/google/uuid"
)

type IImageService interface {
	UploadImage(ctx context.Context, eventID uuid.UUID, file *multipart.FileHeader) (*models.Image, error)
	ListImages(ctx context.Context, eventID uuid.UUID) ([]*models.Image, error)
	DeleteImage(ctx context.Context, eventID, imageID uuid.UUID) error
}
