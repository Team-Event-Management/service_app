package imageservice

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"event_management/internal/models"
	eventrepo "event_management/internal/repositories/event_repositories"
	imagerepo "event_management/internal/repositories/image_repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImageServiceImpl struct {
	eventRepo eventrepo.IEventRepository
	imageRepo imagerepo.IImageRepository
	storageDir string // e.g. "uploads"
}

func NewImageServiceImpl(eventRepo eventrepo.IEventRepository, imageRepo imagerepo.IImageRepository, storageDir string) IImageService {
	return &ImageServiceImpl{
		eventRepo:  eventRepo,
		imageRepo:  imageRepo,
		storageDir: storageDir,
	}
}

func (s *ImageServiceImpl) UploadImage(ctx context.Context, eventID uuid.UUID, file *multipart.FileHeader) (*models.Image, error) {
	if _, err := s.eventRepo.FindById(ctx, eventID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("event not found")
		}
		return nil, err
	}

	if err := os.MkdirAll(s.storageDir, os.ModePerm); err != nil {
		return nil, err
	}
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	fullpath := filepath.Join(s.storageDir, filename)

	dst, err := os.Create(fullpath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	img := &models.Image{
		ID:        uuid.New(),
		ImagePath: fullpath,
	}
	if err := s.imageRepo.Create(ctx, img); err != nil {
		return nil, err
	}

	if err := s.imageRepo.AttachToEvent(ctx, eventID, img.ID); err != nil {
		return nil, err
	}

	return img, nil
}

func (s *ImageServiceImpl) ListImages(ctx context.Context, eventID uuid.UUID) ([]*models.Image, error) {
	if _, err := s.eventRepo.FindById(ctx, eventID); err != nil {
		return nil, err
	}
	return s.imageRepo.ListByEvent(ctx, eventID)
}

func (s *ImageServiceImpl) DeleteImage(ctx context.Context, eventID, imageID uuid.UUID) error {
	if err := s.imageRepo.DetachFromEvent(ctx, eventID, imageID); err != nil {
		return err
	}
	return s.imageRepo.Delete(ctx, imageID)
}
