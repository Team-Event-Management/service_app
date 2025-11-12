package imagerepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IImageRepository interface {
	Create(ctx context.Context, img *models.Image) error
	AttachToEvent(ctx context.Context, eventID, imageID uuid.UUID) error
	ListByEvent(ctx context.Context, eventID uuid.UUID) ([]*models.Image, error)
	DetachFromEvent(ctx context.Context, eventID, imageID uuid.UUID) error
	Delete(ctx context.Context, imageID uuid.UUID) error
}
