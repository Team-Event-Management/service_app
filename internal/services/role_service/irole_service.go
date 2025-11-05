package roleservice

import (
	"context"
	rolerequest "event_management/internal/dto/request/role_request"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IRoleService interface {
	CreateRole(ctx context.Context, req rolerequest.CreateRoleRequest) error
	GetAllRole(ctx context.Context, page, limit int, search string) ([]*models.Role, int, error)
	GetByIdRole(ctx context.Context, roleId uuid.UUID) (*models.Role, error)
	UpdateRole(ctx context.Context, roleId uuid.UUID, req rolerequest.UpdateRoleRequest) error
	DeleteRole(ctx context.Context, roleId uuid.UUID) error
}
