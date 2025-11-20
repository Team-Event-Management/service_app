package standcategoryrepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StandCategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewStandCategoryRepositoryImpl(db *gorm.DB) IStandCategoryRepository {
	return &StandCategoryRepositoryImpl{DB: db}
}

func (r *StandCategoryRepositoryImpl) Create(ctx context.Context, data *models.StandCategory) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

func (r *StandCategoryRepositoryImpl) FindByName(ctx context.Context, name string) (*models.StandCategory, error) {
	var standCategory models.StandCategory

	if err := r.DB.WithContext(ctx).Where("LOWER(name) = LOWER(?)", name).First(&standCategory).Error; err != nil {
		return nil, err
	}

	return &standCategory, nil
}

func (r *StandCategoryRepositoryImpl) FindAll(ctx context.Context, limit, offset int, search string) ([]*models.StandCategory, int, error) {
	var (
		standCategory []*models.StandCategory
		count         int64
	)

	query := r.DB.WithContext(ctx).Model(&models.StandCategory{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&standCategory).Error; err != nil {
		return nil, 0, err
	}

	return standCategory, int(count), nil
}

func (r *StandCategoryRepositoryImpl) FindById(ctx context.Context, standCategoryId uuid.UUID) (*models.StandCategory, error) {
	var standCategory models.StandCategory

	if err := r.DB.WithContext(ctx).Where("id = ?", standCategoryId).First(&standCategory).Error; err != nil {
		return nil, err
	}

	return &standCategory, nil
}

func (r *StandCategoryRepositoryImpl) Update(ctx context.Context, standCategoryId uuid.UUID, data *models.StandCategory) error {
	var existing models.StandCategory

	if err := r.DB.WithContext(ctx).Where("id = ?", standCategoryId).First(&existing).Error; err != nil {
		return err
	}

	existing.Name = data.Name
	existing.Description = data.Description
	return r.DB.WithContext(ctx).Save(&existing).Error
}

func (r *StandCategoryRepositoryImpl) Delete(ctx context.Context, standCategoryId uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&models.StandCategory{}, "id = ?", standCategoryId).Error
}
