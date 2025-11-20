package standrepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IStandRepository interface {
	Create(ctx context.Context, stand *models.Stand) error
	FindByName(ctx context.Context, name string) (*models.Stand, error)
	FindAll(ctx context.Context, limit, offset int, search string) ([]*models.Stand, int, error)
	FindById(ctx context.Context, id uuid.UUID) (*models.Stand, error)
	Update(ctx context.Context, id uuid.UUID, stand *models.Stand) error
	Delete(ctx context.Context, id uuid.UUID) error
}
