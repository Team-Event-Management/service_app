package standcategoryservice

import (
	"context"
	standcategoryrequest "event_management/internal/dto/request/stand_category_request"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IStandCategoryService interface {
	CreateStandCategory(ctx context.Context, req standcategoryrequest.CreateStandCategoryRequest) error
	GetAllStandCategory(ctx context.Context, page, limit int, search string) ([]*models.StandCategory, int, error)
	GetByIdStandCategory(ctx context.Context, standCategoryId uuid.UUID) (*models.StandCategory, error)
	UpdateStandCategory(ctx context.Context, standCategoryId uuid.UUID, req standcategoryrequest.UpdateStandCategoryRequest) error
	DeleteStandCategory(ctx context.Context, standCategoryId uuid.UUID) error
}
