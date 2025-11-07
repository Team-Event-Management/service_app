package adminservice

import (
	"context"
	adminrequest "event_management/internal/dto/request/admin_request"
	"event_management/internal/models"

	"github.com/google/uuid"
)

type IAdminService interface {
	Register(ctx context.Context, req adminrequest.RegisterAdminRequest) error
	Login(ctx context.Context, req adminrequest.LoginAdminRequest) (string, error)
	GetProfile(ctx context.Context, adminId uuid.UUID) (*models.User, error)
	UpdateProfile(ctx context.Context, adminID uuid.UUID, req adminrequest.UpdateProfileRequest) error
	Logout(ctx context.Context, adminID uuid.UUID) error
}