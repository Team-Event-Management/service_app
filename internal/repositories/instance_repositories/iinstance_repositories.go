package instancerepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IInstanceRepository interface {
	Create(ctx context.Context, instance *models.Instance) error
	FindByName(ctx context.Context, name string) (*models.Instance, error)
	FindAll(ctx context.Context, limit, offset int, search string) ([]*models.Instance, int, error)
	FindById(ctx context.Context, id uuid.UUID) (*models.Instance, error)
	Update(ctx context.Context, id uuid.UUID, instance *models.Instance) error
	Delete(ctx context.Context, id uuid.UUID) error
}
