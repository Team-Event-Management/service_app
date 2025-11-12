package imagerepositories

import (
	"context"
	"event_management/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImageRepositoryImpl struct {
	DB *gorm.DB
}

func NewImageRepositoryImpl(db *gorm.DB) IImageRepository {
	return &ImageRepositoryImpl{DB: db}
}

func (r *ImageRepositoryImpl) Create(ctx context.Context, img *models.Image) error {
	return r.DB.WithContext(ctx).Create(img).Error
}

func (r *ImageRepositoryImpl) AttachToEvent(ctx context.Context, eventID, imageID uuid.UUID) error {
	pivot := &models.ImageEvent{IDEvent: eventID, IDImage: imageID}
	return r.DB.WithContext(ctx).Create(pivot).Error
}

func (r *ImageRepositoryImpl) ListByEvent(ctx context.Context, eventID uuid.UUID) ([]*models.Image, error) {
	var images []*models.Image
	if err := r.DB.WithContext(ctx).
		Joins("JOIN image_events ie ON ie.id_image = images.id").
		Where("ie.id_event = ?", eventID).
		Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (r *ImageRepositoryImpl) DetachFromEvent(ctx context.Context, eventID, imageID uuid.UUID) error {
	return r.DB.WithContext(ctx).
		Delete(&models.ImageEvent{}, "id_event = ? AND id_image = ?", eventID, imageID).Error
}

func (r *ImageRepositoryImpl) Delete(ctx context.Context, imageID uuid.UUID) error {
	return r.DB.WithContext(ctx).Delete(&models.Image{}, "id = ?", imageID).Error
}
