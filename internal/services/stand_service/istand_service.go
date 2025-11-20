package standservice

import (
	"context"
	standrequest "event_management/internal/dto/request/stand_request"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IStandService interface {
	CreateStand(ctx context.Context, req standrequest.CreateStandRequest) error
	GetAllStand(ctx context.Context, page, limit int, search string) ([]*models.Stand, int, error)
	GetByIdStand(ctx context.Context, id uuid.UUID) (*models.Stand, error)
	UpdateStand(ctx context.Context, id uuid.UUID, req standrequest.UpdateStandRequest) error
	DeleteStand(ctx context.Context, id uuid.UUID) error
}
