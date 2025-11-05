package rolerepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) Create(ctx context.Context, data *models.Role) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

func (r *RoleRepository) FindByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role

	if err := r.DB.WithContext(ctx).Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepository) FindAll(ctx context.Context, limit, offset int, search string) ([]*models.Role, int, error) {
	var (
		roles  []*models.Role
		count int64
	)

	query := r.DB.WithContext(ctx).Model(&models.Role{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, int(count), nil
}

func (r *RoleRepository) FindById(ctx context.Context, roleId uuid.UUID) (*models.Role, error) {
	var role models.Role

	if err := r.DB.WithContext(ctx).Where("id = ?", roleId).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepository) Update(ctx context.Context, roleId uuid.UUID, data *models.Role) error {
	var existing models.Role
	
	if err := r.DB.WithContext(ctx).Where("id = ?", roleId).First(&existing).Error; err != nil {
		return err
	}

	existing.Name = data.Name
	return r.DB.WithContext(ctx).Save(&existing).Error
}

func (r *RoleRepository) Delete(ctx context.Context, roleId uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&models.Role{}, "id = ?", roleId).Error
}