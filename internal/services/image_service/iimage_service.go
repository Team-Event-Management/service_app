package imageservice

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IImageService interface {
	ConsumeEventImageUpload()
	ListImages(ctx context.Context, eventID uuid.UUID) ([]*models.Image, error)
	DeleteImage(ctx context.Context, eventID, imageID uuid.UUID) error
}