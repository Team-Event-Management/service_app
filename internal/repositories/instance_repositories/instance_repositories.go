package instancerepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InstanceRepositoryImpl struct {
	DB *gorm.DB
}

func NewInstanceRepositoryImpl(db *gorm.DB) IInstanceRepository {
	return &InstanceRepositoryImpl{DB: db}
}

func (r *InstanceRepositoryImpl) Create(ctx context.Context, data *models.Instance) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

func (r *InstanceRepositoryImpl) FindByName(ctx context.Context, name string) (*models.Instance, error) {
	var instance models.Instance

	if err := r.DB.WithContext(ctx).Where("LOWER(name) = LOWER(?)", name).First(&instance).Error; err != nil {
		return nil, err
	}

	return &instance, nil
}

func (r *InstanceRepositoryImpl) FindAll(ctx context.Context, limit, offset int, search string) ([]*models.Instance, int, error) {
	var (
		instances  []*models.Instance
		count int64
	)

	query := r.DB.WithContext(ctx).Model(&models.Instance{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&instances).Error; err != nil {
		return nil, 0, err
	}

	return instances, int(count), nil
}

func (r *InstanceRepositoryImpl) FindById(ctx context.Context, instanceId uuid.UUID) (*models.Instance, error) {
	var instance models.Instance

	if err := r.DB.WithContext(ctx).Where("id = ?", instanceId).First(&instance).Error; err != nil {
		return nil, err
	}

	return &instance, nil
}

func (r *InstanceRepositoryImpl) Update(ctx context.Context, instanceId uuid.UUID, data *models.Instance) error {
	var existing models.Instance
	
	if err := r.DB.WithContext(ctx).Where("id = ?", instanceId).First(&existing).Error; err != nil {
		return err
	}

	existing.Name = data.Name
	existing.Lat = data.Lat
	existing.Lng = data.Lng
	existing.FullAddress = data.FullAddress
	return r.DB.WithContext(ctx).Save(&existing).Error
}

func (r *InstanceRepositoryImpl) Delete(ctx context.Context, instanceId uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&models.Instance{}, "id = ?", instanceId).Error
}