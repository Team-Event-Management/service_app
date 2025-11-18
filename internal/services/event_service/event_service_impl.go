package eventservice

import (
	"context"
	eventrequest "event_management/internal/dto/request/event_request"
	"event_management/internal/models"
	eventrepository "event_management/internal/repositories/event_repositories"
	errorresponse "event_management/pkg/constant/error_response"
	rabbitmq "event_management/pkg/constant/rabbitMq"
	"event_management/pkg/workers/payload"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/google/uuid"
)

type EventServiceImpl struct {
	eventRepo eventrepository.IEventRepository
}

func NewEventServiceImpl(eventRepo eventrepository.IEventRepository) IEventService {
	return &EventServiceImpl{eventRepo: eventRepo}
}

func fileToBytes(fh *multipart.FileHeader) ([]byte, error) {
	f, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

func (s *EventServiceImpl) CreateEvent(ctx context.Context, req eventrequest.CreateEventRequest) error {
	if strings.TrimSpace(req.NameEvent) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Name Event is required", 400)
	}

	event := &models.Event{
		ID:          uuid.New(),
		NameEvent:   req.NameEvent,
		Description: req.Description,
		Status:      req.Status,
		Location:    req.Location,
	}

	// simpan event
	if err := s.eventRepo.Create(ctx, event); err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to create event", 500)
	}

	// kirim tiap gambar ke queue
	for _, imgFile := range req.EventImages {
		binner, err := fileToBytes(imgFile)
		if err != nil || len(binner) == 0 {
			continue
		}

		pay := payload.ImageUploadPayload{
			ID:        event.ID, // cukup UUID, jangan event.ID.String()
			Type:      "many",   // ganti dari "multiple" ke "many" agar sesuai dengan consumer
			FileBytes: binner,
			Folder:    "management_event/event_images",
			Filename:  fmt.Sprintf("event_%s_%s", event.ID, imgFile.Filename),
		}

		_ = rabbitmq.PublishToQueue("", rabbitmq.SendEventImageQueueName, pay)
	}

	return nil
}

func (s *EventServiceImpl) GetAllEvent(ctx context.Context, page, limit int, search string) ([]*models.Event, int, error) {
	offset := (page - 1) * limit
	return s.eventRepo.FindAll(ctx, limit, offset, search)
}

func (s *EventServiceImpl) GetByIdEvent(ctx context.Context, eventId uuid.UUID) (*models.Event, error) {
	return s.eventRepo.FindById(ctx, eventId)
}

func (s *EventServiceImpl) UpdateEvent(ctx context.Context, eventId uuid.UUID, req eventrequest.UpdateEventRequest) error {
	event, err := s.eventRepo.FindById(ctx, eventId)
	if err != nil {
		return err
	}

	if req.NameEvent != "" {
		event.NameEvent = req.NameEvent
	}
	event.Description = req.Description
	event.Status = req.Status
	event.Location = req.Location

	return s.eventRepo.Update(ctx, eventId, event)
}

func (s *EventServiceImpl) DeleteEvent(ctx context.Context, eventId uuid.UUID) error {
	return s.eventRepo.Delete(ctx, eventId)
}
