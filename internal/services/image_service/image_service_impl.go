package imageservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"event_management/configs"
	datasources "event_management/internal/dataSources"
	"event_management/internal/models"
	eventrepo "event_management/internal/repositories/event_repositories"
	imagerepo "event_management/internal/repositories/image_repositories"
	rabbitmq "event_management/pkg/constant/rabbitMq"
	"event_management/pkg/workers/payload"

	"github.com/google/uuid"
)

type ImageServiceImpl struct {
	eventRepo eventrepo.IEventRepository
	imageRepo imagerepo.IImageRepository
	cld       datasources.CloudinaryService
}

func NewImageServiceImpl(eventRepo eventrepo.IEventRepository, imageRepo imagerepo.IImageRepository, cld datasources.CloudinaryService) IImageService {
	return &ImageServiceImpl{
		eventRepo: eventRepo,
		imageRepo: imageRepo,
		cld:       cld,
	}
}

func (s *ImageServiceImpl) ConsumeEventImageUpload() {
	ch := configs.GetRabbitChannel()
	q, _ := ch.QueueDeclare(rabbitmq.SendEventImageQueueName, true, false, false, false, nil)
	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	fmt.Println("📥 Listening for event image uploads...")

	for msg := range msgs {
		var pay payload.ImageUploadPayload
		if err := json.Unmarshal(msg.Body, &pay); err != nil {
			continue
		}

		reader := bytes.NewReader(pay.FileBytes)
		res, err := s.cld.UploadImageBytes(context.Background(), reader, pay.Folder, pay.Filename)
		if err != nil {
			continue
		}

		img := &models.Image{
			ID:        uuid.New(),
			ImagePath: res.URL,
		}

		if err := s.imageRepo.Create(context.Background(), img); err != nil {
			continue
		}

		eventID := pay.ID

		if err := s.imageRepo.AttachToEvent(context.Background(), eventID, img.ID); err != nil {
			continue
		}
	}
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
