package eventrepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IEventRepository interface {
	Create(ctx context.Context, instance *models.Event) error
	FindByName(ctx context.Context, name string) (*models.Event, error)
	FindAll(ctx context.Context, limit, offset int, search string) ([]*models.Event, int, error)
	FindById(ctx context.Context, id uuid.UUID) (*models.Event, error)
	Update(ctx context.Context, id uuid.UUID, instance *models.Event) error
	Delete(ctx context.Context, id uuid.UUID) error
}