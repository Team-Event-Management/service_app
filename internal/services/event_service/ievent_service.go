package eventservice

import (
	"context"
	eventrequest "event_management/internal/dto/request/event_request"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IEventService interface {
	CreateEvent(ctx context.Context, req eventrequest.CreateEventRequest) error
	GetAllEvent(ctx context.Context, page, limit int, search string) ([]*models.Event, int, error)
	GetByIdEvent(ctx context.Context, eventId uuid.UUID) (*models.Event, error)
	UpdateEvent(ctx context.Context, eventId uuid.UUID, req eventrequest.UpdateEventRequest) error
	DeleteEvent(ctx context.Context, eventId uuid.UUID) error
}
