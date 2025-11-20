package standcategoryservice

import (
	"context"
	standcategoryrequest "event_management/internal/dto/request/stand_category_request"
	"event_management/internal/models"
	standcategoryrepositories "event_management/internal/repositories/stand_category_repositories"

	"github.com/google/uuid"
)

type StandCategoryServiceImpl struct {
	standCategoryRepo standcategoryrepositories.IStandCategoryRepository
}

func NewStandCategoryServiceImpl(standCategoryRepo standcategoryrepositories.IStandCategoryRepository) IStandCategoryService {
	return &StandCategoryServiceImpl{standCategoryRepo: standCategoryRepo}
}

func (s *StandCategoryServiceImpl) CreateStandCategory(ctx context.Context, req standcategoryrequest.CreateStandCategoryRequest) error {
	standCategory := &models.StandCategory{
		Name:        req.Name,
		Description: req.Description,
	}
	return s.standCategoryRepo.Create(ctx, standCategory)
}

func (s *StandCategoryServiceImpl) GetAllStandCategory(ctx context.Context, page, limit int, search string) ([]*models.StandCategory, int, error) {
	offset := (page - 1) * limit
	return s.standCategoryRepo.FindAll(ctx, limit, offset, search)
}

func (s *StandCategoryServiceImpl) GetByIdStandCategory(ctx context.Context, standCategoryId uuid.UUID) (*models.StandCategory, error) {
	return s.standCategoryRepo.FindById(ctx, standCategoryId)
}

func (s *StandCategoryServiceImpl) UpdateStandCategory(ctx context.Context, standCategoryId uuid.UUID, req standcategoryrequest.UpdateStandCategoryRequest) error {
	standCategory, err := s.standCategoryRepo.FindById(ctx, standCategoryId)
	if err != nil {
		return err
	}

	standCategory.Name = req.Name
	standCategory.Description = req.Description
	return s.standCategoryRepo.Update(ctx, standCategoryId, standCategory)
}

func (s *StandCategoryServiceImpl) DeleteStandCategory(ctx context.Context, standCategoryId uuid.UUID) error {
	return s.standCategoryRepo.Delete(ctx, standCategoryId)
}
