package standrepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StandRepositoryImpl struct {
	DB *gorm.DB
}

func NewStandRepositoryImpl(db *gorm.DB) IStandRepository {
	return &StandRepositoryImpl{DB: db}
}

func (r *StandRepositoryImpl) Create(ctx context.Context, data *models.Stand) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

func (r *StandRepositoryImpl) FindByName(ctx context.Context, name string) (*models.Stand, error) {
	var stand models.Stand

	if err := r.DB.WithContext(ctx).Where("LOWER(name) = LOWER(?)", name).First(&stand).Error; err != nil {
		return nil, err
	}

	return &stand, nil
}

func (r *StandRepositoryImpl) FindAll(ctx context.Context, limit, offset int, search string) ([]*models.Stand, int, error) {
	var (
		stands []*models.Stand
		count  int64
	)

	query := r.DB.WithContext(ctx).Model(&models.Stand{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&stands).Error; err != nil {
		return nil, 0, err
	}

	return stands, int(count), nil
}

func (r *StandRepositoryImpl) FindById(ctx context.Context, standId uuid.UUID) (*models.Stand, error) {
	var stand models.Stand

	if err := r.DB.WithContext(ctx).Where("id = ?", standId).First(&stand).Error; err != nil {
		return nil, err
	}

	return &stand, nil
}

func (r *StandRepositoryImpl) Update(ctx context.Context, standId uuid.UUID, data *models.Stand) error {
    var existing models.Stand

    if err := r.DB.WithContext(ctx).Where("id = ?", standId).First(&existing).Error; err != nil {
        return err
    }

    existing.StandName = data.StandName
    existing.Lat = data.Lat
    existing.Lng = data.Lng
    existing.Address = data.Address
    existing.StandCategoryID = data.StandCategoryID
	
    return r.DB.WithContext(ctx).Save(&existing).Error
}


func (r *StandRepositoryImpl) Delete(ctx context.Context, standId uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&models.Stand{}, "id = ?", standId).Error
}
