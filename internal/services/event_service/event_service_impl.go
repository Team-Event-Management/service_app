package eventservice

import (
	"context"
	eventrequest "event_management/internal/dto/request/event_request"
	"event_management/internal/models"
	eventrepository "event_management/internal/repositories/event_repositories"
	imageRepository "event_management/internal/repositories/image_repositories"
	imageService "event_management/internal/services/image_service"
	errorresponse "event_management/pkg/constant/error_response"
	"strings"

	"github.com/google/uuid"
)

type EventServiceImpl struct {
	eventRepo eventrepository.IEventRepository
}

func NewEventServiceImpl(eventRepo eventrepository.IEventRepository) IEventService {
	return &EventServiceImpl{eventRepo: eventRepo}
}

func (s *EventServiceImpl) CreateEvent(ctx context.Context, req eventrequest.CreateEventRequest) error {
	imageRepo := imageRepository.NewImageRepositoryImpl(s.eventRepo.(*eventrepository.EventRepositoryImpl).DB)
	imageService := imageService.NewImageServiceImpl(s.eventRepo, imageRepo, "uploads")

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

	createEventErr := s.eventRepo.Create(ctx, event)
	if createEventErr != nil {
		return createEventErr
	}

	// manggil image service buat tiap image
	var uploadErrors []string
	for i := 0; i < len(req.EventImages); i++ {
		// upload image
		_, err := imageService.UploadImage(ctx, event.ID, req.EventImages[i])
		if err != nil {
			uploadErrors = append(uploadErrors, err.Error())
		}
	}

	if len(uploadErrors) > 0 {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Some event_images failed to upload: "+strings.Join(uploadErrors, "; "), 207)
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
