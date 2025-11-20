package standcategoryrepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IStandCategoryRepository interface {
	Create(ctx context.Context, standCategory *models.StandCategory) error
	FindByName(ctx context.Context, name string) (*models.StandCategory, error)
	FindAll(ctx context.Context, limit, offset int, search string) ([]*models.StandCategory, int, error)
	FindById(ctx context.Context, id uuid.UUID) (*models.StandCategory, error)
	Update(ctx context.Context, id uuid.UUID, standCategory *models.StandCategory) error
	Delete(ctx context.Context, id uuid.UUID) error
}
