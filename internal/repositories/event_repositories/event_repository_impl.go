package eventrepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	DB *gorm.DB
}

func NewEventRepositoryImpl(db *gorm.DB) IEventRepository {
	return &EventRepositoryImpl{DB: db}
}

func (r *EventRepositoryImpl) Create(ctx context.Context, data *models.Event) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

func (r *EventRepositoryImpl) FindByName(ctx context.Context, name string) (*models.Event, error) {
	var event models.Event

	if err := r.DB.WithContext(ctx).Where("LOWER(name) = LOWER(?)", name).First(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *EventRepositoryImpl) FindAll(ctx context.Context, limit, offset int, search string) ([]*models.Event, int, error) {
	var (
		events  []*models.Event
		count int64
	)

	query := r.DB.WithContext(ctx).Model(&models.Event{})
	query.Preload("EventImages")

	if search != "" {
		query = query.Where("name_event ILIKE ?", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Preload("EventImages").Limit(limit).Offset(offset).Order("created_at DESC").Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, int(count), nil
}

func (r *EventRepositoryImpl) FindById(ctx context.Context, eventId uuid.UUID) (*models.Event, error) {
	var event models.Event

	if err := r.DB.WithContext(ctx).Preload("EventImages").Where("id = ?", eventId).First(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *EventRepositoryImpl) Update(ctx context.Context, eventId uuid.UUID, data *models.Event) error {
	return r.DB.WithContext(ctx).Model(&models.Event{}).Where("id = ?", eventId).Updates(data).Error
}

func (r *EventRepositoryImpl) Delete(ctx context.Context, eventId uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&models.Event{}, "id = ?", eventId).Error
}

