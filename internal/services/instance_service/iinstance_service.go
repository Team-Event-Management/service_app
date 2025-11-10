package instanceservice

import (
	"context"
	instancerequest "event_management/internal/dto/request/instance_request"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IInstanceService interface {
	CreateInstance(ctx context.Context, req instancerequest.CreateInstanceRequest) error
	GetAllInstance(ctx context.Context, page, limit int, search string) ([]*models.Instance, int, error)
	GetByIdInstance(ctx context.Context, instanceId uuid.UUID) (*models.Instance, error)
	UpdateInstance(ctx context.Context, instanceId uuid.UUID, req instancerequest.UpdateInstanceRequest) error
	DeleteInstance(ctx context.Context, instanceId uuid.UUID) error
}